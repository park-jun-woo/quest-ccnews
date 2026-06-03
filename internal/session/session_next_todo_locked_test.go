//ff:func feature=article type=helper control=sequence
//ff:what NextTODO가 모든 기사가 잠긴 상태일 때 nil을 반환하는지 검증한다(래칫).

package session

import "testing"

func TestNextTODOLocked(t *testing.T) {
	s := &Session{Articles: []*Article{
		{URL: "a", State: PASS},
		{URL: "b", State: DONE},
	}}
	if got := s.NextTODO(); got != nil {
		t.Errorf("NextTODO() = %v, want nil", got)
	}
}
