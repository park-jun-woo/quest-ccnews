//ff:func feature=cli type=helper control=sequence
//ff:what submit SKIPPED 포매터. extract.Apply 신뢰 게이트가 구조화 데이터 없음으로 기사를 SKIPPED로 잠가 앵커 게이트가 미실행됐음을 알린다(재시도 불가).

package cmd

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// printSubmitSkipped reports an article that the trust gate (extract.Apply) just
// locked to SKIPPED before the anchor gate ran — no structured data, so the
// submitted event6 was never evaluated. SKIPPED is terminal (재시도 불가).
func printSubmitSkipped(w io.Writer, a *session.Article) {
	fmt.Fprintf(w, "판정: %s\n", session.SKIPPED)
	fmt.Fprintf(w, "사유: %s\n", a.SkipReason)
	fmt.Fprintf(w, "기사 상태: %s\n", a.State)
	fmt.Fprintln(w, "구조화 데이터가 없어 신뢰 게이트에서 제외 — 앵커 게이트 미실행, 재제출 불가.")
}
