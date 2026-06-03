//ff:func feature=robots type=helper control=sequence
//ff:what SelectGroup이 token-prefix 일치(specific) 그룹을 * 그룹보다 우선 선택하는지 검증한다.

package robots

import "testing"

func TestSelectGroupSpecificBeatsStar(t *testing.T) {
	rs := &Ruleset{Groups: []Group{
		{Agents: []string{"*"}},
		{Agents: []string{"parkjunwoo"}},
	}}
	g := SelectGroup(rs, "parkjunwoo-quest")
	if g == nil || len(g.Agents) == 0 || g.Agents[0] != "parkjunwoo" {
		t.Errorf("expected specific group, got %+v", g)
	}
}
