//ff:func feature=robots type=helper control=sequence
//ff:what Allow가 먼저 나오고 같은 길이 Disallow가 뒤따라도 Evaluate가 Allow를 우선하는지(순서 무관) 검증한다.

package robots

import "testing"

func TestEvaluateTieAllowWinsAllowFirst(t *testing.T) {
	// Allow appears first, then equal-length Disallow; Allow must still win.
	rs := &Ruleset{Groups: []Group{{Agents: []string{"*"}, Rules: []Rule{
		{Allow: true, Pattern: "/page"},
		{Allow: false, Pattern: "/page"},
	}}}}
	d := Evaluate(rs, "parkjunwoo-quest", "/page")
	if !d.Allowed || d.Rule != "Allow: /page" {
		t.Errorf("tie should favor Allow regardless of order, got %+v", d)
	}
}
