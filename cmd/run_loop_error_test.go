//ff:func feature=cli type=helper control=sequence level=error
//ff:what runIngestion이 인제스천 루프(ingestRun)가 에러를 내면 그 에러를 그대로 반환하는지 검증한다(stub seam).

package cmd

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/ingest"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/spf13/cobra"
)

// TestRunIngestion_LoopError stubs ingestRun to fail; runIngestion must return
// that error before reaching the sweep stage.
func TestRunIngestion_LoopError(t *testing.T) {
	dir := t.TempDir()
	// Existing valid session so the load branch succeeds (not os.ErrNotExist).
	s := session.New("ua", "cc-news")
	sessionPath = writeSessionFile(t, dir, s)
	withRunStub(t, dir)

	wantErr := errors.New("boom")
	ingestRun = func(*ingest.Client, *session.Session, ingest.RunOptions, io.Writer) error {
		return wantErr
	}

	cmd := &cobra.Command{}
	cmd.SetOut(new(bytes.Buffer))
	if err := runIngestion(cmd, nil); !errors.Is(err, wantErr) {
		t.Fatalf("err = %v, want %v", err, wantErr)
	}
}
