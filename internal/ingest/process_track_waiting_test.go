//ff:func feature=ingestion type=helper control=sequence
//ff:what ProcessTrack(forward)가 현재 월을 전부 처리했으면 waiting으로 멈추는지 검증한다.

package ingest

import (
	"testing"
	"time"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestProcessTrackForwardWaitingWhenAllProcessed(t *testing.T) {
	obj := "crawl-data/CC-NEWS/2026/06/CC-NEWS-20260615000000-00001.warc.gz"
	srv := ingestServer(t, obj, warcBytes(t, "https://x.com/a"))
	defer srv.Close()

	c := clientTo(srv, t.TempDir())
	s := session.New("ua", "cc-news")
	s.Ingestion.ProcessedWarcs = []string{"CC-NEWS-20260615000000-00001.warc.gz"}
	now := time.Date(2026, 6, 20, 0, 0, 0, 0, time.UTC)

	res, err := c.ProcessTrack(s, Forward, now)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if !res.Stopped || res.State != StateWaiting {
		t.Errorf("res = %+v want stopped/waiting", res)
	}
	if s.Ingestion.Forward.State != StateWaiting {
		t.Errorf("cursor state = %q", s.Ingestion.Forward.State)
	}
}
