//ff:func feature=cli type=helper control=sequence level=error
//ff:what runSubmit가 검증·저장 후 종단 기사를 out으로 sweep할 때 out 경로의 부모가 파일이라 append가 실패하면 그 에러를 반환하는지 검증한다.

package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
)

// TestRunSubmit_SweepError drives runSubmit to a terminal PASS verdict (so the
// article is eligible for the output sweep) but points --out at a path whose
// parent is a regular file, making output.Append's MkdirAll fail. runSubmit must
// surface that sweep error.
func TestRunSubmit_SweepError(t *testing.T) {
	resetSubmitFlags(t)
	cache, file := writeCacheWarc(t, "<html><body>Alice met Bob in Paris on Monday to sign the treaty</body></html>")
	p := writeSessionWith(t, file)
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

	// Create a regular file, then use it as a directory component of --out so
	// MkdirAll (and thus Append) inside Sweep fails.
	blocker := filepath.Join(t.TempDir(), "blocker")
	if err := os.WriteFile(blocker, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	prevOut := submitOut
	submitOut = filepath.Join(blocker, "sub", "out.jsonl")
	t.Cleanup(func() { submitOut = prevOut })

	cmd := &cobra.Command{}
	cmd.SetOut(new(bytes.Buffer))
	if err := runSubmit(cmd, nil); err == nil {
		t.Fatal("want sweep append error")
	}
}
