//ff:func feature=event6 type=helper control=sequence
//ff:what decodeEvent6 (e) JSON 아예 없음. 균형 JSON 객체가 없는 출력에서 ok=false를 돌려주는지(Prepare가 retryable FAIL로 흡수) 검증한다.

package ccnewsquest

import "testing"

func TestDecodeEvent6NoJSON(t *testing.T) {
	raw := []byte("I'm sorry, I cannot extract event6 from this article.")
	if _, ok := decodeEvent6(raw); ok {
		t.Fatal("decodeEvent6(no JSON) ok = true, want false")
	}
}
