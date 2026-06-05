//ff:func feature=ingestion type=helper control=sequence
//ff:what rangeSize가 파싱 불가 URL(NewRequest 실패)에 대해 네트워크 호출 전 (0,false)를 돌려주는지 검증한다.

package ingest

import "testing"

func TestRangeSizeBadURL(t *testing.T) {
	c := NewClient("test-ua/1.0", t.TempDir())
	// A control character in the URL makes http.NewRequest fail.
	n, ok := c.rangeSize("http://example.com/\x7f")
	if ok {
		t.Error("ok = true, want false (bad URL)")
	}
	if n != 0 {
		t.Errorf("size = %d, want 0", n)
	}
}
