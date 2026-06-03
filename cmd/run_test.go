//ff:func feature=cli type=command control=sequence
//ff:what runIngestion이 os.ErrNotExist 외의 세션 로드 에러(깨진 JSON)를 그대로 반환하는지 검증한다.

package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
)

// TestRunIngestionLoadError covers the branch where session.Load fails with an
// error other than os.ErrNotExist (here: malformed JSON), which runIngestion
// returns directly. The success path drives ingest.Run, which performs real
// CC-NEWS network IO and is therefore left as best-effort (not exercised here).
func TestRunIngestionLoadError(t *testing.T) {
	dir := t.TempDir()
	bad := filepath.Join(dir, "session.json")
	if err := os.WriteFile(bad, []byte("{not valid json"), 0o644); err != nil {
		t.Fatal(err)
	}
	sessionPath = bad
	t.Cleanup(func() { sessionPath = "session.json" })

	cmd := &cobra.Command{}
	cmd.SetOut(new(bytes.Buffer))
	if err := runIngestion(cmd, nil); err == nil {
		t.Fatal("want error from malformed session file")
	}
}
