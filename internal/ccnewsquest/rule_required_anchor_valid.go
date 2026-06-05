//ff:func feature=gate type=rule control=sequence
//ff:what required-anchor-valid(Fail). 필수 필드가 존재·위생적이지만 유효앵커가 0개(textmatch.Normalize 후 비지 않은 앵커 없음)면 발동. anchor.checkRequired의 unanchored 분기 — 필수는 앵커 없으면 검증 불가=FAIL.

package ccnewsquest

import (
	"github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"
)

// ruleRequiredAnchorValid fires (FAIL) when a required field that is present and
// hygienic has zero valid anchors (anchor.checkRequired's unanchored branch). It is
// guarded by present+validValue so it never overlaps required-present and never
// touches a nil field; the anchor check uses reins textmatch (Normalize) to count
// valid anchors. A required field cannot be verified with no anchors → FAIL.
func ruleRequiredAnchorValid() gate.Rule {
	return gate.Rule{
		Meta: gate.RuleMeta{ID: "required-anchor-valid", Level: gate.LevelFail,
			Desc: "필수 필드 유효 앵커 0개(검증 불가)"},
		Check: func(ctx gate.Context) (bool, quest.Fact) {
			required, _, ok := event6Of(ctx)
			if !ok {
				return false, quest.Fact{}
			}
			for _, nf := range required {
				if nf.f == nil || !validValue(nf.f.Value) {
					continue // owned by required-present
				}
				if st, _ := checkField(nf.f, ctx.Source); st == statusUnanchored {
					return true, quest.Fact{Where: nf.name,
						Expected: "≥1 유효 앵커", Actual: "유효 앵커 없음"}
				}
			}
			return false, quest.Fact{}
		},
	}
}
