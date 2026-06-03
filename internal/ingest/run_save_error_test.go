//ff:func feature=ingestion type=helper control=sequence
//ff:what Run이 Save 콜백 실패 시 "save session" 에러를 반환하는지 검증한다.

package ingest

import (
	"bytes"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestRunSaveError(t *testing.T) {
	obj := "crawl-data/CC-NEWS/2026/06/CC-NEWS-20260615000000-00001.warc.gz"
	srv := ingestServer(t, obj, warcBytes(t, "https://x.com/a"))
	defer srv.Close()

	c := clientTo(srv, t.TempDir())
	s := session.New("ua", "cc-news")
	opt := RunOptions{
		Tracks: []Track{Forward},
		Now:    time.Date(2026, 6, 20, 0, 0, 0, 0, time.UTC),
		Save:   func() error { return errors.New("disk full") },
	}
	err := Run(c, s, opt, new(bytes.Buffer))
	if err == nil || !strings.Contains(err.Error(), "save session") {
		t.Fatalf("want save error, got %v", err)
	}
}
