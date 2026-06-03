//ff:func feature=cli type=helper control=sequence
//ff:what loadEvent6("-")가 stdin에서 event6 JSON을 읽어 역직렬화하는지 검증한다.

package cmd

import (
	"strings"
	"testing"
)

func TestLoadEvent6_Stdin(t *testing.T) {
	js := `{"what":{"value":"signed","anchors":["sign"]}}`
	ev, err := loadEvent6("-", strings.NewReader(js))
	if err != nil {
		t.Fatalf("loadEvent6 stdin: %v", err)
	}
	if ev.What == nil || ev.What.Value != "signed" {
		t.Errorf("parsed event6 = %+v", ev)
	}
}
