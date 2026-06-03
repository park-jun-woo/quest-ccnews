//ff:func feature=robots type=helper control=sequence
//ff:what 일치하는 specific 그룹이 여러 개일 때 SelectGroup이 첫 그룹을 유지하는지 검증한다.

package robots

import "testing"

func TestSelectGroupFirstSpecificKept(t *testing.T) {
	rs := &Ruleset{Groups: []Group{
		{Agents: []string{"park"}, Rules: []Rule{{Pattern: "/first"}}},
		{Agents: []string{"park"}, Rules: []Rule{{Pattern: "/second"}}},
	}}
	g := SelectGroup(rs, "parkjunwoo")
	if g == nil || g.Rules[0].Pattern != "/first" {
		t.Errorf("expected first matching specific group, got %+v", g)
	}
}
