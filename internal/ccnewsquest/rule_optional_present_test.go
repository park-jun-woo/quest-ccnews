//ff:func feature=gate type=helper control=sequence
//ff:what optional-present 규칙 단위테스트. 선택 필드가 존재하고 value가 플레이스홀더/공허면 발동, nil 선택 필드는 무시(no-fire), 위생적이면 no-fire.

package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/gate"
)

func TestRuleOptionalPresent(t *testing.T) {
	r := ruleOptionalPresent()
	reqOK := func() *session.Event6 {
		return &session.Event6{Who: fld("Alice", "Alice"), What: fld("signed treaty", "sign the treaty")}
	}

	t.Run("non-event6 no-fire", func(t *testing.T) {
		if fired, _ := r.Check(gate.Context{Submission: "nope", Source: testSource}); fired {
			t.Fatal("want no-fire for non-Event6 submission")
		}
	})

	t.Run("placeholder optional value fires", func(t *testing.T) {
		ev := reqOK()
		ev.Where = fld("N/A", "Paris")
		if fired, _ := r.Check(ctxOf(ev, testSource)); !fired {
			t.Fatal("want fired for placeholder optional value")
		}
	})
	t.Run("nil optional ignored", func(t *testing.T) {
		ev := reqOK()
		if fired, _ := r.Check(ctxOf(ev, testSource)); fired {
			t.Fatal("want no-fire when optional fields are nil")
		}
	})
	t.Run("hygienic optional no-fire", func(t *testing.T) {
		ev := reqOK()
		ev.Where = fld("Paris", "Paris")
		if fired, _ := r.Check(ctxOf(ev, testSource)); fired {
			t.Fatal("want no-fire for hygienic optional value")
		}
	})
}
