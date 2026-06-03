//ff:func feature=anchor type=helper control=sequence
//ff:what 필수 필드 앵커가 원문에 없으면 FAIL이고 Reason이 "환각"을 언급하는지 검증한다.

package anchor

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestGate_RequiredHallucinationFail(t *testing.T) {
	ev := &session.Event6{
		Who:  fld("Alice", "Alice"),
		When: fld("Tuesday", "Tuesday"), // not in source
		What: fld("signed treaty", "sign the treaty"),
	}
	res := Gate(ev, gateSource)
	if res.Verdict != FAIL {
		t.Fatalf("Verdict = %s, want FAIL", res.Verdict)
	}
	if !strings.Contains(res.Reason, "환각") {
		t.Errorf("Reason %q should mention hallucination", res.Reason)
	}
}
