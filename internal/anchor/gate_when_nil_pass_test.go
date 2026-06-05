//ff:func feature=anchor type=helper control=sequence
//ff:what when이 nil이어도 필수 who/what이 위생적 value+유효앵커면 PASS인지 검증한다(Phase010: when 선택화 → 생략 가능).

package anchor

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestGate_WhenNilPass(t *testing.T) {
	ev := &session.Event6{
		Who:  fld("Alice", "Alice"),
		What: fld("signed treaty", "sign the treaty"),
		// When omitted: now optional (Phase010), so nil must not FAIL.
	}
	res := Gate(ev, gateSource)
	if res.Verdict != PASS {
		t.Fatalf("Verdict = %s (%s), want PASS", res.Verdict, res.Reason)
	}
	if !ev.Who.Anchored || !ev.What.Anchored {
		t.Error("required who/what should be marked Anchored")
	}
}
