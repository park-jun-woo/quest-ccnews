//ff:func feature=cli type=command control=sequence level=error
//ff:what `ccnews submit --url --event6` 명령. 제출된 event6를 대상 기사 원문에 앵커 게이트로 검증해 PASS/REVIEW면 잠금, FAIL이면 tries++(MaxTries 초과 시 DONE)하고 세션을 저장한다.

package cmd

import (
	"fmt"
	"time"

	"github.com/park-jun-woo/quest-ccnews/internal/anchor"
	"github.com/park-jun-woo/quest-ccnews/internal/extract"
	"github.com/park-jun-woo/quest-ccnews/internal/ingest"
	"github.com/park-jun-woo/quest-ccnews/internal/output"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/spf13/cobra"
)

var (
	submitURL      string
	submitEvent6   string
	submitCacheDir string
	submitOut      string
)

var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "에이전트의 event6 제출을 앵커 게이트로 검증·잠금/재시도",
	Long: `--url 로 지정한 기사에 대해 에이전트가 만든 event6(JSON: --event6 <file> 또는 -=stdin)를
받아, 그 기사 원문을 WARC에서 재독해 앵커 게이트로 검증한다.

  PASS   필수·선택 앵커 전부 원문 substring → state=PASS(잠금)
  REVIEW 필수는 PASS, 선택필드 앵커 0개 존재     → state=REVIEW(잠금)
  FAIL   필수 누락 또는 앵커 환각              → tries++ → TODO 유지(MaxTries 초과 시 DONE)`,
	RunE: runSubmit,
}

// runSubmit verifies a submitted event6 against the target article's source text
// and applies the verdict. IO: session load/save and WARC re-read; the gate
// (anchor.Gate) and the transition (anchor.Apply) are pure.
func runSubmit(cmd *cobra.Command, _ []string) error {
	if submitURL == "" {
		return fmt.Errorf("--url 필수")
	}
	if submitEvent6 == "" {
		return fmt.Errorf("--event6 필수 (파일 경로 또는 - = stdin)")
	}

	s, err := session.Load(sessionPath)
	if err != nil {
		return err
	}
	a, err := s.Find(submitURL)
	if err != nil {
		return err
	}
	if a.State != session.TODO {
		return fmt.Errorf("기사가 TODO가 아님 (현재 %s) — 잠긴 상태는 재제출 불가", a.State)
	}

	ev, err := loadEvent6(submitEvent6, cmd.InOrStdin())
	if err != nil {
		return err
	}

	client := ingest.NewClient(s.UserAgent, submitCacheDir)
	htmlBytes, err := client.ReadBody(a.WARC)
	if err != nil {
		return fmt.Errorf("원문 재독 실패 (%s): %w", a.URL, err)
	}
	res := extract.Parse(htmlBytes)

	verdict := anchor.Gate(ev, res.BodyText)
	anchor.Apply(a, ev, verdict, time.Now().UTC().Format(time.RFC3339))

	if err := s.Save(sessionPath); err != nil {
		return err
	}

	// Sweep terminal, not-yet-emitted articles (uniform across all terminal
	// states) to the JSONL output, then persist the Emitted flags.
	n, err := output.Sweep(s, submitOut)
	if err != nil {
		return err
	}
	if n > 0 {
		if err := s.Save(sessionPath); err != nil {
			return err
		}
	}

	printSubmit(cmd.OutOrStdout(), a, verdict)
	return nil
}

func init() {
	submitCmd.Flags().StringVar(&submitURL, "url", "", "대상 기사 URL (필수)")
	submitCmd.Flags().StringVar(&submitEvent6, "event6", "", "event6 JSON 파일 경로 (- = stdin)")
	submitCmd.Flags().StringVar(&submitCacheDir, "cache-dir", "warc-cache", "다운로드 WARC 캐시 디렉터리")
	submitCmd.Flags().StringVar(&submitOut, "out", defaultOutPath, "종단 기사 JSONL 출력 경로")
	rootCmd.AddCommand(submitCmd)
}
