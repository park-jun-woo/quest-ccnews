//ff:func feature=cli type=helper control=sequence level=error
//ff:what runSubmit가 --event6 파일이 없을 때 loadEvent6 에러를 반환하는지 검증한다.

package cmd

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
)

func TestRunSubmit_LoadEvent6Error(t *testing.T) {
	resetSubmitFlags(t)
	_, file := writeCacheWarc(t, "<html><body>x</body></html>")
	p := writeSessionWith(t, file)
	submitURL = "https://example.com/a"
	submitEvent6 = filepath.Join(t.TempDir(), "missing-ev.json")
	sessionPath = p
	cmd := &cobra.Command{}
	cmd.SetOut(new(bytes.Buffer))
	if err := runSubmit(cmd, nil); err == nil {
		t.Fatal("want loadEvent6 error")
	}
}
