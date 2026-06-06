//ff:func feature=event6 type=helper control=sequence
//ff:what decodeEvent6 (b) ```json 펜스 두른 JSON. 소형 모델의 마크다운 펜스 envelope이 벗겨져 구제 디코드되는지 검증한다.

package ccnewsquest

import "testing"

func TestDecodeEvent6Fenced(t *testing.T) {
	raw := []byte("```json\n" +
		`{"who":{"value":"Acme","anchors":["Acme"]},` +
		`"what":{"value":"recall","anchors":["recall"]}}` +
		"\n```")
	ev, ok := decodeEvent6(raw)
	if !ok {
		t.Fatal("decodeEvent6(fenced JSON) ok = false, want true")
	}
	if ev.Who == nil || ev.Who.Value != "Acme" {
		t.Fatalf("who = %+v, want value Acme", ev.Who)
	}
}
