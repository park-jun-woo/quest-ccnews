//ff:func feature=anchor type=helper control=sequence
//ff:what 필수 필드(who)가 nil이면 FAIL이고 Reason이 그 필드명을 가리키는지 검증한다.

package anchor

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestGate_RequiredNilFail(t *testing.T) {
	ev := &session.Event6{
		When: fld("Monday", "Monday"),
		What: fld("signed treaty", "sign the treaty"),
	}
	res := Gate(ev, gateSource)
	if res.Verdict != FAIL {
		t.Fatalf("Verdict = %s, want FAIL", res.Verdict)
	}
	if !strings.Contains(res.Reason, "who") {
		t.Errorf("Reason %q should name who", res.Reason)
	}
}
