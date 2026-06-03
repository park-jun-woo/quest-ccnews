//ff:func feature=ingestion type=helper control=sequence
//ff:what ProcessTrack(forward)가 WARC를 받아 기사를 추가하고 커서·processed를 래칫하는지 검증한다.

package ingest

import (
	"testing"
	"time"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestProcessTrackForwardAddsArticles(t *testing.T) {
	obj := "crawl-data/CC-NEWS/2026/06/CC-NEWS-20260615000000-00001.warc.gz"
	srv := ingestServer(t, obj, warcBytes(t, "https://x.com/a", "https://y.com/b"))
	defer srv.Close()

	c := clientTo(srv, t.TempDir())
	s := session.New("ua", "cc-news")
	now := time.Date(2026, 6, 20, 0, 0, 0, 0, time.UTC)

	res, err := c.ProcessTrack(s, Forward, now)
	if err != nil {
		t.Fatalf("ProcessTrack error: %v", err)
	}
	if res.Stopped {
		t.Fatal("should not be stopped")
	}
	if res.ArticlesAdd != 2 || len(s.Articles) != 2 {
		t.Errorf("added=%d articles=%d want 2", res.ArticlesAdd, len(s.Articles))
	}
	wantName := "CC-NEWS-20260615000000-00001.warc.gz"
	if res.WarcName != wantName {
		t.Errorf("warcName = %q", res.WarcName)
	}
	// cursor ratcheted
	if s.Ingestion.Forward.Cursor != wantName || s.Ingestion.Forward.State != StateRunning {
		t.Errorf("cursor = %+v", s.Ingestion.Forward)
	}
	if len(s.Ingestion.ProcessedWarcs) != 1 {
		t.Errorf("processed = %v", s.Ingestion.ProcessedWarcs)
	}
}
