//ff:func feature=session type=helper control=sequence
//ff:what Load가 nil Hosts 맵을 빈 맵으로 초기화하는지 검증한다.

package session

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadNilHostsInit(t *testing.T) {
	p := filepath.Join(t.TempDir(), "nohosts.json")
	if err := os.WriteFile(p, []byte(`{"version":1,"user_agent":"a"}`), 0o644); err != nil {
		t.Fatal(err)
	}
	s, err := Load(p)
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}
	if s.Hosts == nil {
		t.Error("Load did not initialize nil Hosts map")
	}
}
