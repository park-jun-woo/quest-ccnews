//ff:func feature=gate type=helper control=sequence
//ff:what optional-unverifiable 규칙 단위테스트. 존재·위생적 선택 필드의 유효앵커가 0개면 발동(REVIEW 레벨), 유효앵커 있으면 no-fire, nil 선택 필드는 no-fire. 규칙 레벨이 LevelReview인지도 확인.

package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/gate"
)

func TestRuleOptionalUnverifiable(t *testing.T) {
	r := ruleOptionalUnverifiable()
	if r.Meta.Level != gate.LevelReview {
		t.Fatalf("Level = %v, want LevelReview", r.Meta.Level)
	}
	reqOK := func() *session.Event6 {
		return &session.Event6{Who: fld("Alice", "Alice"), What: fld("signed treaty", "sign the treaty")}
	}

	t.Run("non-event6 no-fire", func(t *testing.T) {
		if fired, _ := r.Check(gate.Context{Submission: "nope", Source: testSource}); fired {
			t.Fatal("want no-fire for non-Event6 submission")
		}
	})
	t.Run("anchorless present optional fires", func(t *testing.T) {
		ev := reqOK()
		ev.How = fld("met") // present, hygienic value, zero anchors
		if fired, _ := r.Check(ctxOf(ev, testSource)); !fired {
			t.Fatal("want fired for anchorless present optional field")
		}
	})
	t.Run("anchored optional no-fire", func(t *testing.T) {
		ev := reqOK()
		ev.How = fld("met", "met")
		if fired, _ := r.Check(ctxOf(ev, testSource)); fired {
			t.Fatal("want no-fire when optional field is anchored")
		}
	})
	t.Run("nil optional no-fire", func(t *testing.T) {
		ev := reqOK()
		if fired, _ := r.Check(ctxOf(ev, testSource)); fired {
			t.Fatal("want no-fire when optional fields are nil")
		}
	})
}
