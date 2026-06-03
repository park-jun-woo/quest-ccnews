//ff:func feature=robots type=helper control=sequence
//ff:what 더 긴 패턴(Allow)이 짧은 Disallow를 이기는 최장 일치 우선을 Evaluate가 지키는지 검증한다.

package robots

import "testing"

func TestEvaluateLongestMatchWins(t *testing.T) {
	// /a (Disallow) vs /a/b/c (Allow); longest pattern wins → Allow.
	rs := &Ruleset{Groups: []Group{{Agents: []string{"*"}, Rules: []Rule{
		{Allow: false, Pattern: "/a"},
		{Allow: true, Pattern: "/a/b/c"},
	}}}}
	d := Evaluate(rs, "parkjunwoo-quest", "/a/b/c/d")
	if !d.Allowed || d.Rule != "Allow: /a/b/c" {
		t.Errorf("longest match should win, got %+v", d)
	}
}
