//ff:func feature=robots type=helper control=sequence
//ff:what 콕 집은 그룹이 없을 때 SelectGroup이 * 그룹으로 폴백하는지 검증한다.

package robots

import "testing"

func TestSelectGroupFallbackToStar(t *testing.T) {
	rs := &Ruleset{Groups: []Group{
		{Agents: []string{"someotherbot"}},
		{Agents: []string{"*"}},
	}}
	g := SelectGroup(rs, "parkjunwoo-quest")
	if g == nil || g.Agents[0] != "*" {
		t.Errorf("expected star group, got %+v", g)
	}
}
