//ff:func feature=cli type=command control=sequence level=error
//ff:what `ccnews run [--track]` 명령. 세션을 로드(없으면 생성)하고 투트랙 인제스천 루프를 돌려 WARC를 받아 기사를 채운다.

package cmd

import (
	"errors"
	"os"

	"fmt"

	"github.com/park-jun-woo/quest-ccnews/internal/ingest"
	"github.com/park-jun-woo/quest-ccnews/internal/output"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/spf13/cobra"
)

// defaultOutPath is the default JSONL output file (Phase007 열린 결정 확정).
const defaultOutPath = "ccnews-results.jsonl"

// defaultUserAgent matches Phase001 결정 4 (robots UA).
const defaultUserAgent = "parkjunwoo-quest/0.1 (+https://www.parkjunwoo.com)"

var (
	runTrack    string
	runMaxWarcs int
	runCacheDir string
	runOut      string
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "투트랙 WARC 인제스천 루프(다운로드→레코드→기사 추가→커서 전진)",
	Long: `CC-NEWS WARC를 받아 response 레코드를 TODO 기사로 채운다.

  --track forward   최신 덤프만 따라잡고 따라잡으면 waiting
  --track backward  과거로 내려가며 2016-08(exhausted)까지
  --track both      둘 다(기본)

processed_warcs 래칫으로 중복 처리를 막고, 매 WARC 처리 후 세션을 저장하므로
중단해도 커서에서 재개한다.`,
	RunE: runIngestion,
}

// runIngestion loads (or creates) the session, builds the download client, and
// drives the ingestion loop, persisting after every WARC for resumability.
func runIngestion(cmd *cobra.Command, _ []string) error {
	s, err := session.Load(sessionPath)
	if errors.Is(err, os.ErrNotExist) {
		s = session.New(defaultUserAgent, "cc-news")
	} else if err != nil {
		return err
	}

	client := ingest.NewClient(s.UserAgent, runCacheDir)
	opt := ingest.RunOptions{
		Tracks:   ingest.TracksFor(runTrack),
		MaxWarcs: runMaxWarcs,
		Save:     func() error { return s.Save(sessionPath) },
	}
	if err := ingest.Run(client, s, opt, cmd.OutOrStdout()); err != nil {
		return err
	}

	// Sweep terminal, not-yet-emitted articles to the JSONL output, then persist
	// the Emitted flags so the same article is never appended twice.
	n, err := output.Sweep(s, runOut)
	if err != nil {
		return err
	}
	if n > 0 {
		if err := s.Save(sessionPath); err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "emit: +%d records → %s\n", n, runOut)
	}
	return nil
}

func init() {
	runCmd.Flags().StringVar(&runTrack, "track", "both", "처리할 트랙: forward|backward|both")
	runCmd.Flags().IntVar(&runMaxWarcs, "max-warcs", 0, "처리할 WARC 최대 개수(0=무제한)")
	runCmd.Flags().StringVar(&runCacheDir, "cache-dir", "warc-cache", "다운로드 WARC 캐시 디렉터리")
	runCmd.Flags().StringVar(&runOut, "out", defaultOutPath, "종단 기사 JSONL 출력 경로")
	rootCmd.AddCommand(runCmd)
}
