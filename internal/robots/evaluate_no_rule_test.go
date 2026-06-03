//ff:func feature=robots type=helper control=sequence
//ff:what 매칭되는 규칙이 없으면 Evaluate가 default-allow를 반환하는지 검증한다.

package robots

import "testing"

func TestEvaluateNoMatchingRuleDefaultAllow(t *testing.T) {
	rs := &Ruleset{Groups: []Group{{Agents: []string{"*"}, Rules: []Rule{
		{Allow: false, Pattern: "/secret"},
	}}}}
	d := Evaluate(rs, "parkjunwoo-quest", "/public")
	if !d.Allowed || d.Rule != "" {
		t.Errorf("no matching rule → default-allow, got %+v", d)
	}
}
