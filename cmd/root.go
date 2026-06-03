package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// sessionPath: 모든 명령이 공유하는 세션 상태 파일.
var sessionPath string

var rootCmd = &cobra.Command{
	Use:   "ccnews",
	Short: "뉴스 기사에서 육하원칙(event6)을 추출하는 퀘스트 CLI",
	Long: `ccnews — 생성은 AI가, 판정은 기계가.

뉴스 기사 1건 = 1퀘스트. AI가 본문에서 육하원칙(누가·언제·어디서·무엇·어떻게·왜)을
뽑지만, 완료 판정 권한은 AI에게 없다. 게이트가 결정론적으로 검증한다:

  - robots 게이트(호스트당 1회, 캐싱): 크롤링 허용 여부
  - 추출 캐스케이드: JSON-LD/OG(공짜) → 사이트별 셀렉터 템플릿(캐싱) → AI(미스 시)
  - 사실 앵커 게이트: event6의 핵심 사실 토큰이 원문에 실재하는지 대조(환각 차단)

사실만 농축해 공개하므로 저작권 안전(사실은 저작권 비보호). robots를 존중해 수집한다.`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&sessionPath, "session", "session.json", "세션 상태 파일 경로")
}
