//ff:func feature=cli type=helper control=sequence level=error
//ff:what saveAndSweep가 첫 세션 Save가 읽기전용 디렉터리에서 실패할 때 그 에러를 반환하는지 검증한다(루트 등으로 실패하지 않으면 Skip).

package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestSaveAndSweep_SaveError(t *testing.T) {
	roDir := t.TempDir()
	if err := os.Chmod(roDir, 0o500); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chmod(roDir, 0o700) })

	prevSess := sessionPath
	sessionPath = filepath.Join(roDir, "session.json")
	t.Cleanup(func() { sessionPath = prevSess })

	s := session.New("ua", "cc-news")
	if err := saveAndSweep(s); err == nil {
		t.Skip("Save did not fail (filesystem ignores read-only dir, e.g. running as root)")
	}
}
