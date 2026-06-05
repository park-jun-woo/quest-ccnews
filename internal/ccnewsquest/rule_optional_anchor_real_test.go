//ff:func feature=gate type=helper control=sequence
//ff:what optional-anchor-real 규칙 단위테스트. 존재·위생적 선택 필드의 유효앵커가 원문에 없으면(환각) 발동, 전부 substring이면 no-fire, 유효앵커 0개(미검증)는 이 규칙 소관 아님(no-fire).

package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/gate"
)

func TestRuleOptionalAnchorReal(t *testing.T) {
	r := ruleOptionalAnchorReal()
	reqOK := func() *session.Event6 {
		return &session.Event6{Who: fld("Alice", "Alice"), What: fld("signed treaty", "sign the treaty")}
	}

	t.Run("non-event6 no-fire", func(t *testing.T) {
		if fired, _ := r.Check(gate.Context{Submission: "nope", Source: testSource}); fired {
			t.Fatal("want no-fire for non-Event6 submission")
		}
	})

	t.Run("hallucinated optional anchor fires", func(t *testing.T) {
		ev := reqOK()
		ev.Where = fld("Paris", "Atlantis")
		fired, fact := r.Check(ctxOf(ev, testSource))
		if !fired {
			t.Fatal("want fired for hallucinated optional anchor")
		}
		if fact.Actual != "Atlantis" {
			t.Fatalf("Fact.Actual = %q, want %q", fact.Actual, "Atlantis")
		}
	})
	t.Run("all optional anchors real no-fire", func(t *testing.T) {
		ev := reqOK()
		ev.Where = fld("Paris", "Paris")
		if fired, _ := r.Check(ctxOf(ev, testSource)); fired {
			t.Fatal("want no-fire when optional anchors are substrings")
		}
	})
	t.Run("anchorless optional not owned here", func(t *testing.T) {
		ev := reqOK()
		ev.Where = fld("Paris") // no anchors → optional-unverifiable's job
		if fired, _ := r.Check(ctxOf(ev, testSource)); fired {
			t.Fatal("want no-fire: zero-anchor optional is REVIEW, not hallucination")
		}
	})
}
