//ff:func feature=article type=helper control=sequence
//ff:what Find가 없는 URL에 대해 nil과 에러를 반환하는지 검증한다.

package session

import "testing"

func TestFindNotFound(t *testing.T) {
	s := &Session{Articles: []*Article{{URL: "https://example.com/a"}}}

	got, err := s.Find("https://example.com/missing")
	if err == nil {
		t.Fatal("Find expected error, got nil")
	}
	if got != nil {
		t.Errorf("Find = %v, want nil on error", got)
	}
}
