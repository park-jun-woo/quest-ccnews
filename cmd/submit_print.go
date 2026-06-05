//ff:func feature=cli type=helper control=sequence
//ff:what submit 결과 포매터. 게이트 판정(PASS/REVIEW/FAIL)·사유·전이 후 기사 상태/시도수를 찍는다. FAIL이면 재시도/DONE 안내. (SKIPPED는 printSubmitSkipped 별도 파일에서 안내.)

package cmd

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/quest-ccnews/internal/anchor"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// printSubmit reports the gate verdict and the resulting article state after the
// transition has been applied (and the session saved).
func printSubmit(w io.Writer, a *session.Article, res anchor.Result) {
	fmt.Fprintf(w, "판정: %s\n", res.Verdict)
	fmt.Fprintf(w, "사유: %s\n", res.Reason)
	fmt.Fprintf(w, "기사 상태: %s (시도 %d/%d)\n", a.State, a.Tries, session.MaxTries)

	if res.Verdict == anchor.FAIL {
		if a.State == session.DONE {
			fmt.Fprintln(w, "MaxTries 초과 — DONE으로 잠겼습니다(재시도 불가).")
		} else {
			fmt.Fprintln(w, "이번 시도 실패 — TODO 유지. 본문을 다시 보고 앵커를 고쳐 재제출하세요.")
		}
	}
}
