//ff:func feature=output type=helper control=iteration dimension=1
//ff:what boolPtr가 true/false 양쪽에서 입력 값을 가리키는 non-nil 포인터를 반환하는지 검증한다.

package output

import "testing"

func TestBoolPtr(t *testing.T) {
	for _, b := range []bool{true, false} {
		p := boolPtr(b)
		if p == nil || *p != b {
			t.Errorf("boolPtr(%v) = %v", b, p)
		}
	}
}
