//ff:func feature=anchor type=helper control=sequence
//ff:what 앵커 0개인 선택 필드가 둘이면 REVIEW Reason이 첫 번째(where)를 가리키는지 검증한다.

package anchor

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestGate_FirstAnchorlessOptionalNamedInReview(t *testing.T) {
	ev := &session.Event6{
		Who:   fld("Alice", "Alice"),
		When:  fld("Monday", "Monday"),
		What:  fld("signed treaty", "sign the treaty"),
		Where: fld("Paris"), // first anchorless optional → named
		Why:   fld("peace"), // also anchorless
	}
	res := Gate(ev, gateSource)
	if res.Verdict != REVIEW {
		t.Fatalf("Verdict = %s, want REVIEW", res.Verdict)
	}
	if !strings.Contains(res.Reason, "where") {
		t.Errorf("Reason %q should name first anchorless optional (where)", res.Reason)
	}
}
