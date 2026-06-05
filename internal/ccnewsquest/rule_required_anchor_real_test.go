//ff:func feature=gate type=helper control=sequence
//ff:what required-anchor-real 규칙 단위테스트. 필수 필드 유효앵커가 원문에 없으면(환각) 발동하고 Fact.Actual에 환각 앵커 표면형이 실린다. 전부 substring이면 no-fire.

package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/gate"
)

func TestRuleRequiredAnchorReal(t *testing.T) {
	r := ruleRequiredAnchorReal()

	t.Run("non-event6 no-fire", func(t *testing.T) {
		if fired, _ := r.Check(gate.Context{Submission: "nope", Source: testSource}); fired {
			t.Fatal("want no-fire for non-Event6 submission")
		}
	})

	t.Run("hallucinated anchor fires with surface form", func(t *testing.T) {
		ev := &session.Event6{Who: fld("Alice", "Zelda"), What: fld("signed treaty", "sign the treaty")}
		fired, fact := r.Check(ctxOf(ev, testSource))
		if !fired {
			t.Fatal("want fired for hallucinated anchor")
		}
		if fact.Actual != "Zelda" {
			t.Fatalf("Fact.Actual = %q, want %q", fact.Actual, "Zelda")
		}
	})
	t.Run("all anchors real no-fire", func(t *testing.T) {
		ev := &session.Event6{Who: fld("Alice", "Alice"), What: fld("signed treaty", "sign the treaty")}
		if fired, _ := r.Check(ctxOf(ev, testSource)); fired {
			t.Fatal("want no-fire when every anchor is a substring")
		}
	})
	t.Run("nil/invalid required skipped (continue) then later field fires", func(t *testing.T) {
		// who is nil → continue (owned by required-present); what is hallucinated → fires.
		// Exercises the nil/!validValue continue guard before the hallucination check.
		ev := &session.Event6{Who: nil, What: fld("signed treaty", "Nonexistent")}
		fired, fact := r.Check(ctxOf(ev, testSource))
		if !fired {
			t.Fatal("want fired: what's anchor is hallucinated")
		}
		if fact.Actual != "Nonexistent" {
			t.Fatalf("Fact.Actual = %q, want %q", fact.Actual, "Nonexistent")
		}
	})
}
