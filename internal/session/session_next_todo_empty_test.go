//ff:func feature=article type=helper control=sequence
//ff:what NextTODO가 빈 기사 목록에서 nil을 반환하는지 검증한다.

package session

import "testing"

func TestNextTODOEmpty(t *testing.T) {
	s := &Session{Articles: nil}
	if got := s.NextTODO(); got != nil {
		t.Errorf("NextTODO() = %v, want nil", got)
	}
}
