//ff:func feature=robots type=helper control=sequence
//ff:what 일치하는 specific 그룹이 여러 개일 때 SelectGroup이 둘 다 병합해 모든 Rules를 포함하는지 검증한다.

package robots

import "testing"

func TestSelectGroupSpecificMerged(t *testing.T) {
	rs := &Ruleset{Groups: []Group{
		{Agents: []string{"park"}, Rules: []Rule{{Pattern: "/first"}}},
		{Agents: []string{"park"}, Rules: []Rule{{Pattern: "/second"}}},
	}}
	g := SelectGroup(rs, "parkjunwoo")
	if g == nil || len(g.Rules) != 2 {
		t.Fatalf("expected both specific groups merged, got %+v", g)
	}
	if g.Rules[0].Pattern != "/first" || g.Rules[1].Pattern != "/second" {
		t.Errorf("expected /first then /second in declaration order, got %+v", g.Rules)
	}
}
