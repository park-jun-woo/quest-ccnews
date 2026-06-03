//ff:func feature=robots type=helper control=selection
//ff:what 최장 일치로 고른 규칙(없으면 nil)을 Decision으로 변환한다. nil이면 default-allow, 아니면 "Verb: pattern" 감사 문자열. 순수 함수.

package robots

import "fmt"

// decisionFor turns the winning rule (nil when none matched) into a Decision:
// nil means default-allow, otherwise the rule's Allow flag plus an audit string
// "Allow: <pattern>" or "Disallow: <pattern>". Pure — no IO.
func decisionFor(best *Rule) Decision {
	switch {
	case best == nil:
		return Decision{Allowed: true}
	case best.Allow:
		return Decision{Allowed: true, Rule: fmt.Sprintf("Allow: %s", best.Pattern)}
	default:
		return Decision{Allowed: false, Rule: fmt.Sprintf("Disallow: %s", best.Pattern)}
	}
}
