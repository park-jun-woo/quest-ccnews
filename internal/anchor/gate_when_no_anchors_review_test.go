//ff:func feature=anchor type=helper control=sequence
//ff:what 필수 who/what은 PASS이고 선택 when이 present인데 유효앵커 0개면 REVIEW이고 Reason이 when을 가리키는지 검증한다(Phase010: when 선택화).

package anchor

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestGate_WhenNoAnchorsReview(t *testing.T) {
	ev := &session.Event6{
		Who:  fld("Alice", "Alice"),
		What: fld("signed treaty", "sign the treaty"),
		When: fld("2026-06-04"), // present, no anchors → unverifiable
	}
	res := Gate(ev, gateSource)
	if res.Verdict != REVIEW {
		t.Fatalf("Verdict = %s (%s), want REVIEW", res.Verdict, res.Reason)
	}
	if !strings.Contains(res.Reason, "when") {
		t.Errorf("Reason %q should name the anchorless optional field when", res.Reason)
	}
}
