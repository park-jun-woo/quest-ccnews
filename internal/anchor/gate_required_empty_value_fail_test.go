//ff:func feature=anchor type=helper control=sequence
//ff:what 필수 필드의 Value가 비면 FAIL이고 Reason이 "값 없음"을 언급하는지 검증한다.

package anchor

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestGate_RequiredEmptyValueFail(t *testing.T) {
	ev := &session.Event6{
		Who:  fld("", "Alice"),
		When: fld("Monday", "Monday"),
		What: fld("signed treaty", "sign the treaty"),
	}
	res := Gate(ev, gateSource)
	if res.Verdict != FAIL {
		t.Fatalf("Verdict = %s, want FAIL", res.Verdict)
	}
	if !strings.Contains(res.Reason, "값 없음") {
		t.Errorf("Reason %q should mention empty value", res.Reason)
	}
}
