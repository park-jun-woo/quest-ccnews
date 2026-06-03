//ff:func feature=robots type=helper control=sequence
//ff:what 잘못된 호스트로 요청 생성 자체가 실패하면 Fetch가 nil rec/rs와 에러를 반환하는지 검증한다.

package robots

import "testing"

func TestFetchBadRequestURL(t *testing.T) {
	// A host containing a control character makes http.NewRequest fail before
	// any network call, so Fetch returns the construction error.
	c := NewClient("parkjunwoo-quest/0.1")
	rec, rs, err := c.Fetch("exa\x7fmple.com")
	if err == nil {
		t.Fatalf("expected error from invalid request URL")
	}
	if rec != nil || rs != nil {
		t.Errorf("on request build error rec and rs should be nil, got rec=%+v rs=%+v", rec, rs)
	}
}
