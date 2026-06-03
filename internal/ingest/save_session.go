//ff:func feature=ingestion type=helper control=sequence
//ff:what 인제스천 루프의 세션 저장 콜백을 호출한다(설정돼 있으면). 에러는 래핑해 반환한다.

package ingest

import "fmt"

// saveSession invokes the run option's persistence callback when one is set,
// wrapping any failure. Extracted from the Run loop so the per-track step stays
// within the iteration depth limit.
func saveSession(opt RunOptions) error {
	if opt.Save == nil {
		return nil
	}
	if err := opt.Save(); err != nil {
		return fmt.Errorf("save session: %w", err)
	}
	return nil
}
