//ff:func feature=cli type=helper control=sequence
//ff:what runIngestion이 세션 파일이 없으면 새 세션을 만들고, ingestRun(stub)이 종단 기사를 채우면 sweep이 1건 emit하며 "emit: +1" 안내를 찍는지 검증한다.

package cmd

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/ingest"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/spf13/cobra"
)

// TestRunIngestion_NewSessionEmit drives the success path with a stubbed
// ingestRun (no network): the session file is absent so a fresh session is
// created, the stub adds one PASS (terminal) article, and the post-loop sweep
// must emit exactly one JSONL record and report it.
func TestRunIngestion_NewSessionEmit(t *testing.T) {
	dir := t.TempDir()
	sessionPath = filepath.Join(dir, "session.json") // does not exist yet
	withRunStub(t, dir)

	ingestRun = func(_ *ingest.Client, s *session.Session, _ ingest.RunOptions, _ io.Writer) error {
		s.Articles = append(s.Articles, &session.Article{
			URL: "https://e.com/a", Host: "e.com", State: session.PASS,
		})
		return nil
	}

	cmd := &cobra.Command{}
	var out bytes.Buffer
	cmd.SetOut(&out)
	if err := runIngestion(cmd, nil); err != nil {
		t.Fatalf("runIngestion: %v", err)
	}
	if !strings.Contains(out.String(), "emit: +1") {
		t.Errorf("output = %q, want emit: +1", out.String())
	}
	data, err := os.ReadFile(runOut)
	if err != nil || !strings.Contains(string(data), "https://e.com/a") {
		t.Errorf("out file = %q, err=%v", string(data), err)
	}
}
