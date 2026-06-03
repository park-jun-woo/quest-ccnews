//ff:func feature=session type=helper control=sequence
//ff:what Save가 부모가 디렉터리가 아닌 잘못된 경로에 대해 에러를 반환하는지 검증한다.

package session

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSaveError(t *testing.T) {
	s := New("agent", "src")
	// Writing into a path whose parent is a file (not a directory) fails.
	dir := t.TempDir()
	notADir := filepath.Join(dir, "file")
	if err := os.WriteFile(notADir, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	badPath := filepath.Join(notADir, "session.json")
	if err := s.Save(badPath); err == nil {
		t.Errorf("Save(%q) expected error, got nil", badPath)
	}
}
