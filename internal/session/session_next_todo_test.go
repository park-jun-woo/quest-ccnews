//ff:func feature=article type=helper control=sequence
//ff:what NextTODO가 첫 TODO 기사를 고르고 그 앞의 잠긴 기사는 건너뛰는지 검증한다.

package session

import "testing"

func TestNextTODO(t *testing.T) {
	s := &Session{Articles: []*Article{
		{URL: "a", State: PASS},
		{URL: "b", State: TODO},
		{URL: "c", State: TODO},
	}}
	got := s.NextTODO()
	if got == nil {
		t.Fatal("NextTODO() = nil, want article b")
	}
	if got.URL != "b" {
		t.Errorf("NextTODO().URL = %q, want %q", got.URL, "b")
	}
}
