//ff:func feature=gate type=rule control=sequence
//ff:what required-anchor-real(Fail). 필수 필드가 존재·위생적이고 유효앵커 중 하나라도 원문 substring이 아니면(환각) 발동. anchor.checkRequired의 hallucination 분기 — textmatch.MissingTokens로 첫 환각 앵커를 Fact에 싣는다.

package ccnewsquest

import (
	"github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"
)

// ruleRequiredAnchorReal fires (FAIL) when a required field that is present and
// hygienic has a hallucinated anchor — a valid anchor that is not a substring of the
// source (anchor.checkRequired's hallucination branch). checkField performs the
// per-anchor scan via reins textmatch.Contains (matching anchor.Gate) and returns the
// first offender, named in the Fact. Guarded by present+validValue so it never
// overlaps required-present.
func ruleRequiredAnchorReal() gate.Rule {
	return gate.Rule{
		Meta: gate.RuleMeta{ID: "required-anchor-real", Level: gate.LevelFail,
			Desc: "필수 필드 앵커가 원문에 없음(환각)"},
		Check: func(ctx gate.Context) (bool, quest.Fact) {
			required, _, ok := event6Of(ctx)
			if !ok {
				return false, quest.Fact{}
			}
			for _, nf := range required {
				if nf.f == nil || !validValue(nf.f.Value) {
					continue // owned by required-present
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
