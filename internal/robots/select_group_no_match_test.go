//ff:func feature=robots type=helper control=sequence
//ff:what specific도 * 그룹도 없으면 SelectGroup이 nil을 반환하는지 검증한다.

package robots

import "testing"

func TestSelectGroupNoMatchReturnsNil(t *testing.T) {
	rs := &Ruleset{Groups: []Group{
		{Agents: []string{"someotherbot"}},
	}}
	if g := SelectGroup(rs, "parkjunwoo-quest"); g != nil {
		t.Errorf("expected nil, got %+v", g)
	}
}
