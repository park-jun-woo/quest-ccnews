//ff:func feature=event6 type=helper control=sequence
//ff:what decodeEvent6 (c) 앞뒤 산문 섞인 JSON. 펜스 밖 산문이 붙어도 첫 균형 JSON 객체만 추출돼 디코드되는지 검증한다.

package ccnewsquest

import "testing"

func TestDecodeEvent6Prose(t *testing.T) {
	raw := []byte("Sure, here is the event6 you asked for:\n" +
		`{"who":{"value":"Acme","anchors":["Acme"]},` +
		`"what":{"value":"recall","anchors":["recall"]}}` +
		"\nLet me know if you need anything else.")
	ev, ok := decodeEvent6(raw)
	if !ok {
		t.Fatal("decodeEvent6(prose-wrapped JSON) ok = false, want true")
	}
	if ev.Who == nil || ev.Who.Value != "Acme" {
		t.Fatalf("who = %+v, want value Acme", ev.Who)
	}
}
