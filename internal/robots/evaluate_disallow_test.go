//ff:func feature=robots type=helper control=sequence
//ff:what Disallow 패턴에 매칭되는 path를 Evaluate가 거부하고 매칭 룰 문자열을 반환하는지 검증한다.

package robots

import "testing"

func TestEvaluateDisallowMatch(t *testing.T) {
	rs := &Ruleset{Groups: []Group{{Agents: []string{"*"}, Rules: []Rule{
		{Allow: false, Pattern: "/secret"},
	}}}}
	d := Evaluate(rs, "parkjunwoo-quest", "/secret/page")
	if d.Allowed {
		t.Errorf("expected disallow, got %+v", d)
	}
	if d.Rule != "Disallow: /secret" {
		t.Errorf("rule = %q, want %q", d.Rule, "Disallow: /secret")
	}
}
