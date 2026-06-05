//ff:func feature=gate type=rule control=sequence
//ff:what optional-anchor-real(Fail). 선택 필드가 존재·위생적이고 유효앵커 중 하나라도 원문 substring이 아니면(환각) 발동. anchor.checkOptional의 hallucination 분기 — 선택 필드 환각도 FAIL.

package ccnewsquest

import (
	"github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"
)

// ruleOptionalAnchorReal fires (FAIL) when a present, hygienic optional field has a
// hallucinated anchor — a valid anchor absent from the source (anchor.checkOptional's
// hallucination branch). Guarded by present+validValue so it never overlaps
// optional-present. checkField (reins textmatch.Contains) names the first offender.
func ruleOptionalAnchorReal() gate.Rule {
	return gate.Rule{
		Meta: gate.RuleMeta{ID: "optional-anchor-real", Level: gate.LevelFail,
			Desc: "선택 필드(존재) 앵커가 원문에 없음(환각)"},
		Check: func(ctx gate.Context) (bool, quest.Fact) {
			_, optional, ok := event6Of(ctx)
			if !ok {
				return false, quest.Fact{}
			}
			for _, nf := range optional {
				if nf.f == nil || !validValue(nf.f.Value) {
					continue // nil ignored; invalid owned by optional-present
				}
				if st, bad := checkField(nf.f, ctx.Source); st == statusHallucination {
					return true, quest.Fact{Where: nf.name,
						Expected: "앵커가 원문 substring", Actual: bad}
				}
			}
			return false, quest.Fact{}
		},
	}
}
