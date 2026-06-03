//ff:func feature=robots type=helper control=sequence
//ff:what 와일드카드+앵커 패턴(/*.pdf$)이 .pdf만 막고 .html은 막지 않는지 Evaluate로 검증한다.

package robots

import "testing"

func TestEvaluateWildcard(t *testing.T) {
	rs := &Ruleset{Groups: []Group{{Agents: []string{"*"}, Rules: []Rule{
		{Allow: false, Pattern: "/*.pdf$"},
	}}}}
	if d := Evaluate(rs, "parkjunwoo-quest", "/docs/report.pdf"); d.Allowed {
		t.Errorf("/*.pdf$ should block report.pdf, got %+v", d)
	}
	if d := Evaluate(rs, "parkjunwoo-quest", "/docs/report.html"); !d.Allowed {
		t.Errorf("/*.pdf$ should not block .html, got %+v", d)
	}
}
