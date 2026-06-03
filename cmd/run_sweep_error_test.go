//ff:func feature=cli type=helper control=sequence level=error
//ff:what runIngestion이 종단 기사 sweep 중 out 경로의 부모가 파일이라 append가 실패하면 그 에러를 반환하는지 검증한다(stub seam).

package cmd

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/ingest"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/spf13/cobra"
)

// TestRunIngestion_SweepError stubs ingestRun to add a terminal article, then
// points --out under a regular file so output.Append's MkdirAll fails inside the
// sweep. runIngestion must surface that error.
func TestRunIngestion_SweepError(t *testing.T) {
	dir := t.TempDir()
	s := session.New("ua", "cc-news")
	sessionPath = writeSessionFile(t, dir, s)
	withRunStub(t, dir)

	blocker := filepath.Join(dir, "blocker")
	if err := os.WriteFile(blocker, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	runOut = filepath.Join(blocker, "sub", "out.jsonl")

	ingestRun = func(_ *ingest.Client, s *session.Session, _ ingest.RunOptions, _ io.Writer) error {
		s.Articles = append(s.Articles, &session.Article{
			URL: "https://e.com/a", Host: "e.com", State: session.PASS,
		})
		return nil
	}

	cmd := &cobra.Command{}
	cmd.SetOut(new(bytes.Buffer))
	if err := runIngestion(cmd, nil); err == nil {
		t.Fatal("want sweep append error")
	}
}
