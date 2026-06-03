//ff:func feature=anchor type=helper control=sequence
//ff:what 필수 필드에 앵커가 0개면 FAIL이고 Reason이 "검증 불가"를 언급하는지 검증한다.

package anchor

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestGate_RequiredNoAnchorsFail(t *testing.T) {
	ev := &session.Event6{
		Who:  fld("Alice"), // no anchors
		When: fld("Monday", "Monday"),
		What: fld("signed treaty", "sign the treaty"),
	}
	res := Gate(ev, gateSource)
	if res.Verdict != FAIL {
		t.Fatalf("Verdict = %s, want FAIL", res.Verdict)
	}
	if !strings.Contains(res.Reason, "검증 불가") {
		t.Errorf("Reason %q should mention unanchored required field", res.Reason)
	}
}
