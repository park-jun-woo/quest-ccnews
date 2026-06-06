//ff:func feature=ingestion type=command control=sequence level=error
//ff:what `run` 명령의 본체(G1 스트리밍 인제스천). reins 세션을 로드(없으면 빈 세션)하고 Meta에서 인제스천 스크래치를 복원한 뒤, 기존 ingest.Run 투트랙 루프를 돌린다. cacheDir를 절대경로로 정규화(Phase013 B)해 ingest 클라이언트와 Meta 기록에 쓴다. ingest의 매-WARC Save 콜백을 브리지(scratch→reins Item TODO 시드 + Meta 커서/호스트/cache_dir 보존 + 세션 저장)로 연결해 중단해도 reins 세션 커서에서 재개한다. robots는 시드 시 fetch하지 않고 pick-time(Prepare)으로 지연한다(Phase013 A). 종단 기사 export는 reins submit 자동방출/export 소관(G5).
package runcmd

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/quest-ccnews/internal/ingest"
	"github.com/park-jun-woo/reins/pkg/quest"
	"github.com/spf13/cobra"
)

// ingestRun is the ingestion loop, indirected through a package var so tests can
// substitute a no-network stub. Defaults to the real ingest.Run.
var ingestRun = ingest.Run

// run loads (or creates) the reins session, restores the ingestion scratch from
// Meta, drives the two-track ingest loop, and bridges newly scanned articles into
// reins Items after every WARC (persisting cursor/hosts/cache_dir to Meta for
// resumability). Robots is no longer fetched at seed time — every article is seeded
// TODO and the per-host robots decision happens at pick time in Prepare (Phase013 A).
func (o *options) run(cmd *cobra.Command, _ []string) error {
	s, err := quest.Load(o.sessionPath())
	if os.IsNotExist(err) {
		s = quest.New()
	} else if err != nil {
		return err
	}

	scratch, err := restoreScratch(s, o.userAgent)
	if err != nil {
		return err
	}
	// Normalize the cache dir to an absolute path so submit/next re-read from the
	// same place no matter the CWD (Phase013 B). Fall back to the raw value if the
	// CWD cannot be resolved.
	cacheDir := o.cacheDir
	if abs, err := filepath.Abs(cacheDir); err == nil {
		cacheDir = abs
	}
	w := cmd.OutOrStdout()

	save := func() error { return o.flush(scratch, s, cacheDir, w) }
	client := ingest.NewClient(scratch.UserAgent, cacheDir)
	opt := ingest.RunOptions{
		Tracks:   ingest.TracksFor(o.track),
		MaxWarcs: o.maxWarcs,
		Save:     save,
	}
	if err := ingestRun(client, scratch, opt, w); err != nil {
		return err
	}
	// Final flush in case the last step left unbridged articles.
	return o.flush(scratch, s, cacheDir, w)
}
