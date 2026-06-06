//ff:func feature=event6 type=helper control=sequence
//ff:what decodeEvent6 (d) 값에 이스케이프 따옴표·중괄호 포함. 문자열 리터럴 내 } 와 \" 가 깊이 카운트를 깨지 않는지(스캐너 견고성) 검증한다.

package ccnewsquest

import "testing"

func TestDecodeEvent6Escaped(t *testing.T) {
	// who.value contains both an escaped quote and a literal brace; if the scanner
	// did not track string/escape state it would either stop the string early or
	// miscount depth and truncate the object.
	raw := []byte(`prefix {"who":{"value":"say \"hi\" {now}","anchors":["x"]},` +
		`"what":{"value":"y","anchors":["y"]}} trailing`)
	ev, ok := decodeEvent6(raw)
	if !ok {
		t.Fatal("decodeEvent6(escaped quotes/braces) ok = false, want true")
	}
	if ev.Who == nil || ev.Who.Value != `say "hi" {now}` {
		t.Fatalf("who.value = %q, want %q", func() string {
			if ev.Who == nil {
				return "<nil>"
			}
			return ev.Who.Value
		}(), `say "hi" {now}`)
	}
}
