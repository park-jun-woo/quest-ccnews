//ff:func feature=robots type=helper control=sequence
//ff:what SelectGroup이 product token을 대소문자 무시하고 agent와 매칭하는지 검증한다.

package robots

import "testing"

func TestSelectGroupCaseInsensitive(t *testing.T) {
	rs := &Ruleset{Groups: []Group{
		{Agents: []string{"parkjunwoo"}},
	}}
	g := SelectGroup(rs, "ParkJunWoo-Quest")
	if g == nil {
		t.Errorf("expected case-insensitive match")
	}
}
