//ff:func feature=event6 type=helper control=sequence
//ff:what decodeEvent6 (a) 순수 JSON. 펜스·산문 없는 정상 event6 JSON이 그대로 디코드되는지(ok=true, 필드 적재) 검증한다.

package ccnewsquest

import "testing"

func TestDecodeEvent6Pure(t *testing.T) {
	raw := []byte(`{"who":{"value":"Acme","anchors":["Acme"]},` +
		`"what":{"value":"recall","anchors":["recall"]}}`)
	ev, ok := decodeEvent6(raw)
	if !ok {
		t.Fatal("decodeEvent6(pure JSON) ok = false, want true")
	}
	if ev.Who == nil || ev.Who.Value != "Acme" {
		t.Fatalf("who = %+v, want value Acme", ev.Who)
	}
	if ev.What == nil || ev.What.Value != "recall" {
		t.Fatalf("what = %+v, want value recall", ev.What)
	}
}
