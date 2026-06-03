//ff:func feature=robots type=helper control=sequence
//ff:what 더 긴 패턴이 Disallow일 때 짧은 Allow를 이기는 최장 일치 우선을 Evaluate가 지키는지 검증한다.

package robots

import "testing"

func TestEvaluateLongestMatchDisallowWins(t *testing.T) {
	rs := &Ruleset{Groups: []Group{{Agents: []string{"*"}, Rules: []Rule{
		{Allow: true, Pattern: "/a"},
		{Allow: false, Pattern: "/a/b/c"},
	}}}}
	d := Evaluate(rs, "parkjunwoo-quest", "/a/b/c/d")
	if d.Allowed || d.Rule != "Disallow: /a/b/c" {
		t.Errorf("longest disallow should win, got %+v", d)
	}
}
