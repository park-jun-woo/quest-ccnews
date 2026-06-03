//ff:func feature=anchor type=helper control=sequence
//ff:what present 선택 필드의 앵커가 공백뿐이면(["  "]) 유효앵커 0개로 REVIEW이고 Reason이 "유효앵커 0개"를 언급하는지 검증한다(Phase009 L0).

package anchor

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestGate_OptionalWhitespaceAnchorReview(t *testing.T) {
	ev := &session.Event6{
		Who:   fld("Alice", "Alice"),
		When:  fld("Monday", "Monday"),
		What:  fld("signed treaty", "sign the treaty"),
		Where: fld("Paris", "  "), // whitespace-only anchor → no valid anchors
	}
	res := Gate(ev, gateSource)
	if res.Verdict != REVIEW {
		t.Fatalf("Verdict = %s, want REVIEW (whitespace anchor → structurally unverifiable)", res.Verdict)
	}
	if !strings.Contains(res.Reason, "유효앵커 0개") {
		t.Errorf("Reason %q should mention zero valid anchors", res.Reason)
	}
}
