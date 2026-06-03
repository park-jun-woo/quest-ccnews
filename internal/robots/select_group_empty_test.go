//ff:func feature=robots type=helper control=sequence
//ff:what 빈 룰셋에 대해 SelectGroup이 nil을 반환하는지 검증한다.

package robots

import "testing"

func TestSelectGroupEmptyRuleset(t *testing.T) {
	if g := SelectGroup(&Ruleset{}, "anybody"); g != nil {
		t.Errorf("empty ruleset should return nil, got %+v", g)
	}
}
