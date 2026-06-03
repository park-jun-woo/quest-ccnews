//ff:func feature=anchor type=helper control=sequence
//ff:what 필수 who/when/what만 있고 앵커가 전부 원문 substring이면 PASS이고 세 필드 Anchored가 채워지는지 검증한다.

package anchor

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestGate_RequiredOnlyPass(t *testing.T) {
	ev := &session.Event6{
		Who:  fld("Alice", "Alice"),
		When: fld("Monday", "Monday"),
		What: fld("signed treaty", "sign the treaty"),
	}
	res := Gate(ev, gateSource)
	if res.Verdict != PASS {
		t.Fatalf("Verdict = %s (%s), want PASS", res.Verdict, res.Reason)
	}
	if !ev.Who.Anchored || !ev.When.Anchored || !ev.What.Anchored {
		t.Error("required fields should be marked Anchored")
	}
}
