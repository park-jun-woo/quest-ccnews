//ff:func feature=cli type=helper control=sequence level=error
//ff:what runSubmit가 검증 후 세션 Save가 읽기전용 디렉터리에서 실패할 때 에러를 반환하는지 검증한다(루트 등으로 실패하지 않으면 Skip).

package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/spf13/cobra"
)

func TestRunSubmit_SaveError(t *testing.T) {
	resetSubmitFlags(t)
	cache, file := writeCacheWarc(t, submitStructuredHTML(
		"Alice met Bob in Paris on Monday to sign the treaty. "+
			"The two leaders shook hands before the cameras and pledged to deepen cooperation "+
			"across trade, security and climate over the coming year, aides said afterward."))
	// Place the session file inside a directory we make read-only so the
	// post-verdict Save() (submit.go:73) fails on write.
	roDir := t.TempDir()
	p := filepath.Join(roDir, "session.json")
	s := session.New("ua", "cc-news")
	s.Articles = []*session.Article{{
		URL:   "https://example.com/a",
		Host:  "example.com",
		State: session.TODO,
		WARC:  &session.WARCLoc{File: file, Offset: 0},
	}}
	if err := s.Save(p); err != nil {
		t.Fatal(err)
	}
	if err := os.Chmod(p, 0o400); err != nil {
		t.Fatal(err)
	}
	if err := os.Chmod(roDir, 0o500); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chmod(roDir, 0o700); _ = os.Chmod(p, 0o600) })

	evPath := writeEvent6File(t, `{
		"who":{"value":"Alice","anchors":["Alice"]},
		"when":{"value":"Monday","anchors":["Monday"]},
		"what":{"value":"signed treaty","anchors":["sign the treaty"]}
	}`)
	submitURL = "https://example.com/a"
	submitEvent6 = evPath
	sessionPath = p
	prev := submitCacheDir
	submitCacheDir = cache
	t.Cleanup(func() { submitCacheDir = prev })

	cmd := &cobra.Command{}
	cmd.SetOut(new(bytes.Buffer))
	if err := runSubmit(cmd, nil); err == nil {
		t.Skip("Save did not fail (filesystem ignores read-only dir, e.g. running as root)")
	}
}
