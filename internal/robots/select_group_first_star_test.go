//ff:func feature=robots type=helper control=sequence
//ff:what * 그룹이 여러 개일 때 SelectGroup이 첫 * 그룹을 유지하는지 검증한다.

package robots

import "testing"

func TestSelectGroupFirstStarKept(t *testing.T) {
	rs := &Ruleset{Groups: []Group{
		{Agents: []string{"*"}, Rules: []Rule{{Pattern: "/first"}}},
		{Agents: []string{"*"}, Rules: []Rule{{Pattern: "/second"}}},
	}}
	g := SelectGroup(rs, "nobody")
	if g == nil || len(g.Rules) == 0 || g.Rules[0].Pattern != "/first" {
		t.Errorf("expected first star group, got %+v", g)
	}
}
