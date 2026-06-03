//ff:func feature=cli type=command control=sequence level=error
//ff:what `ccnews next` 명령. 다음 TODO 기사 1건을 골라 WARC에서 원문 본문을 재독·추출해 앵커 대상 텍스트와 함께 에이전트용 event6 작성 프롬프트를 출력한다.

package cmd

import (
	"fmt"

	"github.com/park-jun-woo/quest-ccnews/internal/extract"
	"github.com/park-jun-woo/quest-ccnews/internal/ingest"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/spf13/cobra"
)

var nextCacheDir string

var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "다음 TODO 기사 + 원문 본문 + event6 작성 프롬프트 출력",
	Long: `다음 처리할 TODO 기사 1건을 골라, WARC 로케이터로 원문을 재독(본문 미저장)하고
앵커 대상 본문 텍스트를 추출해 출력한다. 에이전트는 이 본문에서 육하원칙(event6)을
만들어 ccnews submit 으로 제출한다 — value는 영어, anchors는 원문 표면형.`,
	RunE: runNext,
}

// runNext loads the session, picks the next TODO article, re-reads its body from
// the WARC (bodies are never persisted), extracts the anchor-target text, and
// prints it with the event6 authoring prompt. IO (session load, WARC re-read);
// the extraction itself is the pure extract.Parse.
func runNext(cmd *cobra.Command, _ []string) error {
	s, err := session.Load(sessionPath)
	if err != nil {
		return err
	}
	a := s.NextTODO()
	if a == nil {
		fmt.Fprintln(cmd.OutOrStdout(), "처리할 TODO 기사가 없습니다.")
		return nil
	}

	client := ingest.NewClient(s.UserAgent, nextCacheDir)
	htmlBytes, err := client.ReadBody(a.WARC)
	if err != nil {
		return fmt.Errorf("원문 재독 실패 (%s): %w", a.URL, err)
	}
	res := extract.Parse(htmlBytes)

	printNext(cmd.OutOrStdout(), a, res.BodyText)
	return nil
}

func init() {
	nextCmd.Flags().StringVar(&nextCacheDir, "cache-dir", "warc-cache", "다운로드 WARC 캐시 디렉터리")
	rootCmd.AddCommand(nextCmd)
}
