//ff:func feature=gate type=helper control=sequence
//ff:what required-anchor-valid 규칙 단위테스트. 필수 필드가 위생적 value지만 유효앵커 0개(앵커 없음/공백 앵커뿐)면 발동, 유효앵커 있으면 no-fire. nil/플레이스홀더는 required-present 소관이라 no-fire.

package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/gate"
)

func TestRuleRequiredAnchorValid(t *testing.T) {
	r := ruleRequiredAnchorValid()

	t.Run("non-event6 no-fire", func(t *testing.T) {
		if fired, _ := r.Check(gate.Context{Submission: "nope", Source: testSource}); fired {
			t.Fatal("want no-fire for non-Event6 submission")
		}
	})

	t.Run("zero valid anchors fires", func(t *testing.T) {
		ev := &session.Event6{Who: fld("Alice"), What: fld("signed treaty", "sign the treaty")}
		if fired, _ := r.Check(ctxOf(ev, testSource)); !fired {
			t.Fatal("want fired for required field with no anchors")
		}
	})
	t.Run("whitespace-only anchor fires", func(t *testing.T) {
		ev := &session.Event6{Who: fld("Alice", "   "), What: fld("signed treaty", "sign the treaty")}
		if fired, _ := r.Check(ctxOf(ev, testSource)); !fired {
			t.Fatal("want fired: whitespace anchor is not a valid anchor")
		}
	})
	t.Run("valid anchor no-fire", func(t *testing.T) {
		ev := &session.Event6{Who: fld("Alice", "Alice"), What: fld("signed treaty", "sign the treaty")}
		if fired, _ := r.Check(ctxOf(ev, testSource)); fired {
			t.Fatal("want no-fire when a valid anchor is present")
		}
	})
	t.Run("nil owned by required-present", func(t *testing.T) {
		ev := &session.Event6{Who: nil, What: fld("signed treaty", "sign the treaty")}
		if fired, _ := r.Check(ctxOf(ev, testSource)); fired {
			t.Fatal("want no-fire for nil field (required-present owns it)")
		}
	})
}
