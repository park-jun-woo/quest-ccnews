//ff:func feature=article type=helper control=sequence
//ff:what Find가 존재하는 URL로 올바른 기사를 찾는지 검증한다.

package session

import "testing"

func TestFindFound(t *testing.T) {
	a1 := &Article{URL: "https://example.com/a"}
	a2 := &Article{URL: "https://example.com/b"}
	s := &Session{Articles: []*Article{a1, a2}}

	got, err := s.Find("https://example.com/b")
	if err != nil {
		t.Fatalf("Find unexpected error: %v", err)
	}
	if got != a2 {
		t.Errorf("Find = %v, want %v", got, a2)
	}
}
