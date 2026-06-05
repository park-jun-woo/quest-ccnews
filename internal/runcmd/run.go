//ff:func feature=ingestion type=command control=sequence level=error
//ff:what `run` 명령의 본체(G1 스트리밍 인제스천). reins 세션을 로드(없으면 빈 세션)하고 Meta에서 인제스천 스크래치를 복원한 뒤, 기존 ingest.Run 투트랙 루프를 돌린다. ingest의 매-WARC Save 콜백을 브리지(scratch→reins Item 시드 + Meta 커서/호스트 보존 + robots BLOCKED 시드 + 세션 저장)로 연결해 중단해도 reins 세션 커서에서 재개한다. 종단 기사 export는 reins submit 자동방출/export 소관(G5).

package runcmd

import (
	"os"
	"time"

	"github.com/park-jun-woo/quest-ccnews/internal/ingest"
	"github.com/park-jun-woo/reins/pkg/quest"
	"github.com/spf13/cobra"
)

// ingestRun is the ingestion loop, indirected through a package var so tests can
// substitute a no-network stub. Defaults to the real ingest.Run.
var ingestRun = ingest.Run

// run loads (or creates) the reins session, restores the ingestion scratch from
// Meta, drives the two-track ingest loop, and bridges newly scanned articles into
// reins Items after every WARC (persisting cursor/hosts to Meta for resumability).
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
	guard := newRobotsGuard(scratch.UserAgent, scratch.Hosts)
	now := time.Now().UTC().Format(time.RFC3339)
	w := cmd.OutOrStdout()

	save := func() error { return o.flush(scratch, s, guard, now, w) }
	client := ingest.NewClient(scratch.UserAgent, o.cacheDir)
	opt := ingest.RunOptions{
		Tracks:   ingest.TracksFor(o.track),
		MaxWarcs: o.maxWarcs,
		Save:     save,
	}
	if err := ingestRun(client, scratch, opt, w); err != nil {
		return err
	}
	// Final flush in case the last step left unbridged articles.
	return o.flush(scratch, s, guard, now, w)
}
