//ff:func feature=cli type=helper control=sequence level=error
//ff:what runSubmit가 세션 파일이 없을 때 로드 에러를 반환하는지 검증한다.

package cmd

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
)

func TestRunSubmit_LoadError(t *testing.T) {
	resetSubmitFlags(t)
	submitURL = "https://example.com/a"
	submitEvent6 = "ev.json"
	sessionPath = filepath.Join(t.TempDir(), "missing.json")
	cmd := &cobra.Command{}
	cmd.SetOut(new(bytes.Buffer))
	if err := runSubmit(cmd, nil); err == nil {
		t.Fatal("want session load error")
	}
}
