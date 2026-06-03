//ff:func feature=cli type=helper control=sequence
//ff:what loadEvent6가 파일 경로에서 event6 JSON을 읽어 session.Event6로 역직렬화하는지 검증한다.

package cmd

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadEvent6_File(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "ev.json")
	js := `{"who":{"value":"Alice","anchors":["Alice"]}}`
	if err := os.WriteFile(p, []byte(js), 0o644); err != nil {
		t.Fatal(err)
	}
	ev, err := loadEvent6(p, nil)
	if err != nil {
		t.Fatalf("loadEvent6: %v", err)
	}
	if ev.Who == nil || ev.Who.Value != "Alice" || len(ev.Who.Anchors) != 1 {
		t.Errorf("parsed event6 = %+v", ev)
	}
}
