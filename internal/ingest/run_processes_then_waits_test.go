//ff:func feature=ingestion type=helper control=sequence
//ff:what Run이 첫 턴에 WARC를 처리하고 다음 턴에 받을 게 없어 waiting으로 멈추는지 검증한다.

package ingest

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestRunProcessesThenWaits(t *testing.T) {
	obj := "crawl-data/CC-NEWS/2026/06/CC-NEWS-20260615000000-00001.warc.gz"
	srv := ingestServer(t, obj, warcBytes(t, "https://x.com/a"))
	defer srv.Close()

	c := clientTo(srv, t.TempDir())
	s := session.New("ua", "cc-news")
	var out bytes.Buffer
	opt := RunOptions{
		Tracks: []Track{Forward},
		Now:    time.Date(2026, 6, 20, 0, 0, 0, 0, time.UTC),
	}
	if err := Run(c, s, opt, &out); err != nil {
		t.Fatalf("Run error: %v", err)
	}
	// first turn processes the WARC, second turn finds nothing → waiting
	if len(s.Articles) != 1 {
		t.Errorf("articles = %d want 1", len(s.Articles))
	}
	o := out.String()
	if !strings.Contains(o, "+1 articles") || !strings.Contains(o, "stopped: waiting") {
		t.Errorf("unexpected output: %q", o)
	}
}
