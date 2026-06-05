//ff:func feature=gate type=helper control=sequence
//ff:what required-present 규칙 단위테스트. nil 필수 필드·플레이스홀더 value·빈 value면 발동(true), 위생적 value면 no-fire. 비-Event6 제출이면 no-fire. 네트워크 없이 Context 직접 구성.

package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/gate"
)

func TestRuleRequiredPresent(t *testing.T) {
	r := ruleRequiredPresent()

	t.Run("nil who fires", func(t *testing.T) {
		ev := &session.Event6{Who: nil, What: fld("signed treaty", "sign the treaty")}
		if fired, _ := r.Check(ctxOf(ev, testSource)); !fired {
			t.Fatal("want fired for nil required field")
		}
	})
	t.Run("placeholder value fires", func(t *testing.T) {
		ev := &session.Event6{Who: fld("Subject", "Alice"), What: fld("signed treaty", "sign the treaty")}
		if fired, _ := r.Check(ctxOf(ev, testSource)); !fired {
			t.Fatal("want fired for placeholder value")
		}
	})
	t.Run("hygienic required no-fire", func(t *testing.T) {
		ev := &session.Event6{Who: fld("Alice", "Alice"), What: fld("signed treaty", "sign the treaty")}
		if fired, _ := r.Check(ctxOf(ev, testSource)); fired {
			t.Fatal("want no-fire for present hygienic required fields")
		}
	})
	t.Run("non-event6 no-fire", func(t *testing.T) {
		if fired, _ := r.Check(gate.Context{Submission: "nope", Source: testSource}); fired {
			t.Fatal("want no-fire for non-Event6 submission")
		}
	})
}
