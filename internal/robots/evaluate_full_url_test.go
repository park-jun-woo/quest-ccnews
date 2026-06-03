//ff:func feature=robots type=helper control=sequence
//ff:what Evaluate가 전체 URL을 받아 path로 정규화한 뒤 규칙에 매칭하는지 검증한다.

package robots

import "testing"

func TestEvaluateAcceptsFullURL(t *testing.T) {
	rs := &Ruleset{Groups: []Group{{Agents: []string{"*"}, Rules: []Rule{
		{Allow: false, Pattern: "/secret"},
	}}}}
	d := Evaluate(rs, "parkjunwoo-quest", "https://example.com/secret/page")
	if d.Allowed {
		t.Errorf("full URL path should be normalized and matched, got %+v", d)
	}
}
