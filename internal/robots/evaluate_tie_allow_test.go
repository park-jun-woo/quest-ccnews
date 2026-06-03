//ff:func feature=robots type=helper control=sequence
//ff:what 같은 길이의 Allow/Disallow가 모두 매칭되면 Evaluate가 Allow를 우선하는지 검증한다.

package robots

import "testing"

func TestEvaluateTieAllowWins(t *testing.T) {
	// Equal length patterns both match → Allow beats Disallow.
	rs := &Ruleset{Groups: []Group{{Agents: []string{"*"}, Rules: []Rule{
		{Allow: false, Pattern: "/page"},
		{Allow: true, Pattern: "/page"},
	}}}}
	d := Evaluate(rs, "parkjunwoo-quest", "/page")
	if !d.Allowed || d.Rule != "Allow: /page" {
		t.Errorf("tie should favor Allow, got %+v", d)
	}
}
