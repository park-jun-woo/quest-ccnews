//ff:func feature=cli type=helper control=sequence
//ff:what runIngestion이 종단 기사가 없어 sweep이 0건이면 emit 안내 없이 정상 종료하는지 검증한다(stub seam, n==0 분기).

package cmd

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/ingest"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/spf13/cobra"
)

// TestRunIngestion_NoEmit stubs ingestRun to add only a TODO (non-terminal)
// article, so the sweep writes nothing (n==0): runIngestion returns nil without
// an emit line and without creating the out file.
func TestRunIngestion_NoEmit(t *testing.T) {
	dir := t.TempDir()
	s := session.New("ua", "cc-news")
	sessionPath = writeSessionFile(t, dir, s)
	withRunStub(t, dir)

	ingestRun = func(_ *ingest.Client, s *session.Session, _ ingest.RunOptions, _ io.Writer) error {
		s.Articles = append(s.Articles, &session.Article{
			URL: "https://e.com/todo", Host: "e.com", State: session.TODO,
		})
		return nil
	}

	cmd := &cobra.Command{}
	var out bytes.Buffer
	cmd.SetOut(&out)
	if err := runIngestion(cmd, nil); err != nil {
		t.Fatalf("runIngestion: %v", err)
	}
	if strings.Contains(out.String(), "emit:") {
		t.Errorf("output = %q, want no emit line", out.String())
	}
	if _, err := os.Stat(runOut); !os.IsNotExist(err) {
		t.Errorf("out file should not be created when nothing emitted")
	}
}
