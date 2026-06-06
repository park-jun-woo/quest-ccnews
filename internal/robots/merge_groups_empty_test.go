//ff:func feature=robots type=helper control=sequence
//ff:what mergeGroups에 빈 입력(nil·빈 슬라이스)을 주면 nil을 반환하는지 검증한다.

package robots

import "testing"

func TestMergeGroupsEmptyReturnsNil(t *testing.T) {
	if g := mergeGroups(nil); g != nil {
		t.Errorf("mergeGroups(nil) = %+v, want nil", g)
	}
	if g := mergeGroups([]*Group{}); g != nil {
		t.Errorf("mergeGroups([]) = %+v, want nil", g)
	}
}
