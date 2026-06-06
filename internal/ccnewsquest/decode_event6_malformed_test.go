//ff:func feature=event6 type=helper control=sequence
//ff:what decodeEvent6 (e) 균형 잡힌 객체이지만 타입이 안 맞아 Unmarshal이 실패하는 경로. firstJSONObject는 통과시키되 json.Unmarshal이 거부 → (zero,false) 분기 검증.

package ccnewsquest

import "testing"

func TestDecodeEvent6Malformed(t *testing.T) {
	// The braces are balanced so firstJSONObject returns this slice, but "who" is a
	// bare string where Event6.Who is a *Field (object). json.Unmarshal therefore
	// fails, exercising the decode's error branch which must yield (zero, false).
	raw := []byte(`{"who":"not-an-object"}`)
	ev, ok := decodeEvent6(raw)
	if ok {
		t.Fatalf("decodeEvent6(type-mismatch) ok = true, want false (ev=%+v)", ev)
	}
	if ev.Who != nil {
		t.Fatalf("expected zero Event6 on failure, got who=%+v", ev.Who)
	}
}
