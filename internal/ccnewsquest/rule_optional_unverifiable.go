//ff:func feature=gate type=rule control=sequence
//ff:what optional-unverifiable(Review). 선택 필드가 존재·위생적이지만 유효앵커가 0개면 발동(구조적 미검증, 사람 확인 필요). anchor.checkOptional의 unanchored 분기 → anchor.Gate의 REVIEW. 레벨집계가 필수 FAIL이 이걸 이기게 해 anchor.Gate 선후를 보존.

package ccnewsquest

import (
	"github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"
)

// ruleOptionalUnverifiable fires (REVIEW) when a present, hygienic optional field has
// zero valid anchors — structurally unverifiable, a human must confirm
// (anchor.checkOptional's unanchored branch, which anchor.Gate turns into REVIEW).
// Being LevelReview, it loses to any fired required-* Fail rule in gate.Evaluate's
// level aggregation, which reproduces anchor.Gate's "required FAIL returns before the
// optional REVIEW" ordering verdict-for-verdict. Guarded by present+validValue.
func ruleOptionalUnverifiable() gate.Rule {
	return gate.Rule{
		Meta: gate.RuleMeta{ID: "optional-unverifiable", Level: gate.LevelReview,
			Desc: "선택 필드(존재) 유효 앵커 0개 — 구조적 미검증(사람 확인 필요)"},
		Check: func(ctx gate.Context) (bool, quest.Fact) {
			_, optional, ok := event6Of(ctx)
			if !ok {
				return false, quest.Fact{}
			}
			for _, nf := range optional {
				if nf.f == nil || !validValue(nf.f.Value) {
					continue // nil ignored; invalid owned by optional-present
				}
				if st, _ := checkField(nf.f, ctx.Source); st == statusUnanchored {
					return true, quest.Fact{Where: nf.name,
						Expected: "≥1 유효 앵커", Actual: "유효 앵커 없음(구조적 미검증)"}
				}
			}
			return false, quest.Fact{}
		},
	}
}
