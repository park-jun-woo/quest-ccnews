//ff:func feature=anchor type=helper control=sequence
//ff:what 필수+선택 6필드의 앵커가 전부 원문 substring이면 PASS이고 선택 필드 Anchored도 채워지는지 검증한다.

package anchor

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestGate_RequiredAndOptionalPass(t *testing.T) {
	ev := &session.Event6{
		Who:   fld("Alice", "Alice"),
		When:  fld("Monday", "Monday"),
		What:  fld("signed treaty", "sign the treaty"),
		Where: fld("Paris", "Paris"),
		How:   fld("met", "met"),
		Why:   fld("peace", "peace mattered"),
	}
	res := Gate(ev, gateSource)
	if res.Verdict != PASS {
		t.Fatalf("Verdict = %s (%s), want PASS", res.Verdict, res.Reason)
	}
	if !ev.Where.Anchored || !ev.How.Anchored || !ev.Why.Anchored {
		t.Error("optional fields should be marked Anchored")
	}
}
