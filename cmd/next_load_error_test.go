//ff:func feature=cli type=helper control=sequence level=error
//ff:what runNext가 세션 파일이 없을 때 에러를 반환하는지 검증한다.

package cmd

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
)

func TestRunNext_LoadError(t *testing.T) {
	sessionPath = filepath.Join(t.TempDir(), "missing.json")
	t.Cleanup(func() { sessionPath = "session.json" })
	cmd := &cobra.Command{}
	cmd.SetOut(new(bytes.Buffer))
	if err := runNext(cmd, nil); err == nil {
		t.Fatal("want error from missing session")
	}
}
