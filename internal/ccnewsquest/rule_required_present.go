//ff:func feature=gate type=rule control=sequence
//ff:what required-present(Fail). 필수 필드(who/what)가 nil(부재)이거나 value가 무효(빈·룬<2·플레이스홀더)면 발동. anchor.checkRequired의 nil 분기 + !validValue 분기를 합친 것 — "부재" 변명 불가.

package ccnewsquest

import (
	"github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"
)

// ruleRequiredPresent fires (FAIL) when a required field (who/what) is absent (nil)
// or carries an intrinsically invalid value (empty, <2 runes, or a placeholder —
// Phase009 L3). It merges anchor.checkRequired's nil-check and !validValue branches:
// both are "the required fact is not really present" failures the gate cannot let
// the agent excuse as "부재".
func ruleRequiredPresent() gate.Rule {
	return gate.Rule{
		Meta: gate.RuleMeta{ID: "required-present", Level: gate.LevelFail,
			Desc: "필수 필드(who/what) 누락 또는 플레이스홀더/공허 value"},
		Check: func(ctx gate.Context) (bool, quest.Fact) {
			required, _, ok := event6Of(ctx)
			if !ok {
				return false, quest.Fact{}
			}
			for _, nf := range required {
				if nf.f == nil {
					return true, quest.Fact{Where: nf.name,
						Expected: "필수 필드 존재", Actual: "누락(값 없음)"}
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
