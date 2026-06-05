//ff:func feature=gate type=rule control=sequence
//ff:what optional-present(Fail). 선택 필드(when/where/how/why)가 존재(non-nil)하나 value가 무효(빈·룬<2·플레이스홀더)면 발동. nil 선택 필드는 무시. anchor.checkOptional의 !validValue 분기.

package ccnewsquest

import (
	"github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"
)

// ruleOptionalPresent fires (FAIL) when a present optional field (when/where/how/why)
// carries an intrinsically invalid value (empty, <2 runes, or a placeholder). A nil
// optional field is ignored (absent is allowed). This is anchor.checkOptional's
// !validValue branch — checked before the anchor test, exactly as the original.
func ruleOptionalPresent() gate.Rule {
	return gate.Rule{
		Meta: gate.RuleMeta{ID: "optional-present", Level: gate.LevelFail,
			Desc: "선택 필드(존재) 플레이스홀더/공허 value"},
		Check: func(ctx gate.Context) (bool, quest.Fact) {
			_, optional, ok := event6Of(ctx)
			if !ok {
				return false, quest.Fact{}
			}
			for _, nf := range optional {
				if nf.f == nil {
					continue
				}
				if !validValue(nf.f.Value) {
					return true, quest.Fact{Where: nf.name,
						Expected: "위생적 value(비어있지 않음·룬≥2·비플레이스홀더)", Actual: nf.f.Value}
				}
			}
			return false, quest.Fact{}
		},
	}
}
