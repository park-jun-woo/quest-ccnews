//ff:func feature=anchor type=helper control=sequence
//ff:what 필수는 PASS이고 present 선택 필드(why)에 앵커가 0개면 REVIEW이고 Reason이 그 필드명을 가리키는지 검증한다.

package anchor

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestGate_OptionalNoAnchorsReview(t *testing.T) {
	ev := &session.Event6{
		Who:  fld("Alice", "Alice"),
		When: fld("Monday", "Monday"),
		What: fld("signed treaty", "sign the treaty"),
		Why:  fld("peace"), // present, no anchors
	}
	res := Gate(ev, gateSource)
	if res.Verdict != REVIEW {
		t.Fatalf("Verdict = %s (%s), want REVIEW", res.Verdict, res.Reason)
	}
	if !strings.Contains(res.Reason, "why") {
		t.Errorf("Reason %q should name the anchorless optional field", res.Reason)
	}
}
