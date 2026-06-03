//ff:func feature=ingestion type=helper control=sequence
//ff:what ProcessTrack(backward)가 최古 월을 전부 처리했으면 exhausted로 멈추는지 검증한다.

package ingest

import (
	"testing"
	"time"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestProcessTrackBackwardExhausted(t *testing.T) {
	obj := "crawl-data/CC-NEWS/2016/08/CC-NEWS-20160815000000-00001.warc.gz"
	srv := ingestServer(t, obj, warcBytes(t, "https://x.com/a"))
	defer srv.Close()

	c := clientTo(srv, t.TempDir())
	s := session.New("ua", "cc-news")
	// pretend we're already at the earliest month with it processed
	s.Ingestion.Backward = &session.Cursor{Cursor: "CC-NEWS-20160815000000-00001.warc.gz", State: StateRunning}
	s.Ingestion.ProcessedWarcs = []string{"CC-NEWS-20160815000000-00001.warc.gz"}
	now := time.Date(2026, 6, 20, 0, 0, 0, 0, time.UTC)

	res, err := c.ProcessTrack(s, Backward, now)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if !res.Stopped || res.State != StateExhausted {
		t.Errorf("res = %+v want stopped/exhausted", res)
	}
}
