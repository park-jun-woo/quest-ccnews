//ff:func feature=session type=helper control=sequence
//ff:what Load가 깨진 JSON 파일에 대해 에러를 반환하는지 검증한다.

package session

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadInvalidJSON(t *testing.T) {
	p := filepath.Join(t.TempDir(), "bad.json")
	if err := os.WriteFile(p, []byte("{not json"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := Load(p); err == nil {
		t.Error("Load of invalid json expected error, got nil")
	}
}
