//ff:func feature=extract type=helper control=sequence
//ff:what setIfEmpty가 빈 대상엔 첫 값을 쓰고 이미 값이 있으면 덮어쓰지 않는지 검증한다.

package extract

import "testing"

func TestSetIfEmpty(t *testing.T) {
	var s string
	setIfEmpty(&s, "first")
	if s != "first" {
		t.Fatalf("got %q", s)
	}
	setIfEmpty(&s, "second")
	if s != "first" {
		t.Fatalf("setIfEmpty overwrote: %q", s)
	}
}
