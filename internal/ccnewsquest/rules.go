//ff:func feature=gate type=helper control=sequence
//ff:what Definition.Rules. 앵커 규칙 6개(필수 3: present/anchor-valid/anchor-real + 선택 3: present/anchor-real/unverifiable)의 카탈로그를 반환한다. gate.Evaluate가 이 평평한 카탈로그를 레벨집계(필수 FAIL>선택 REVIEW)해 anchor.Gate와 동일 verdict를 낸다. ccnews는 진짜 선후가드가 없어 그래프 불요.

package ccnewsquest

import "github.com/park-jun-woo/reins/pkg/gate"

// Rules returns ccnews's anchor-gate violation catalog: three required rules
// (required-present / -anchor-valid / -anchor-real) and three optional rules
// (optional-present / -anchor-real / -unverifiable). reins gate.Evaluate aggregates
// them by Level — any fired required/optional Fail rule makes the verdict FAIL,
// otherwise a fired optional-unverifiable Review makes it REVIEW, otherwise PASS —
// which is verdict-equivalent to anchor.Gate. ccnews has no genuine precedence
// guards, so the flat catalog path suffices (no defeat graph).
func (ccnewsDef) Rules() []gate.Rule {
	return []gate.Rule{
		ruleRequiredPresent(),
		ruleRequiredAnchorValid(),
		ruleRequiredAnchorReal(),
		ruleOptionalPresent(),
		ruleOptionalAnchorReal(),
		ruleOptionalUnverifiable(),
	}
}
