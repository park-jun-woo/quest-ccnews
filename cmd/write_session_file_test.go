//ff:func feature=cli type=helper control=sequence
//ff:what 테스트 헬퍼. 세션을 dir/session.json에 저장하고 그 경로를 돌려준다(기존-세션 분기용).

package cmd

import (
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// writeSessionFile saves s into dir/session.json and returns the path, so a
// runIngestion test can take the existing-session (non os.ErrNotExist) branch.
func writeSessionFile(t *testing.T, dir string, s *session.Session) string {
	t.Helper()
	p := filepath.Join(dir, "session.json")
	if err := s.Save(p); err != nil {
		t.Fatal(err)
	}
	return p
}
