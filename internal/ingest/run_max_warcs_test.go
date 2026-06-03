//ff:func feature=ingestion type=helper control=sequence
//ff:what Run이 MaxWarcs=1일 때 WARC 한 개만 처리하고 멈추는지 검증한다.

package ingest

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestRunMaxWarcsLimit(t *testing.T) {
	// Two distinct WARCs available in the month; MaxWarcs=1 stops after one.
	objs := []string{
		"crawl-data/CC-NEWS/2026/06/CC-NEWS-20260615000000-00001.warc.gz",
		"crawl-data/CC-NEWS/2026/06/CC-NEWS-20260615000000-00002.warc.gz",
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "warc.paths.gz") {
			w.Write(gzBytes(t, strings.Join(objs, "\n")+"\n"))
			return
		}
		w.Write(warcBytes(t, "https://x.com/a"))
	}))
	defer srv.Close()

	c := clientTo(srv, t.TempDir())
	s := session.New("ua", "cc-news")
	var out bytes.Buffer
	opt := RunOptions{
		Tracks:   []Track{Forward},
		MaxWarcs: 1,
		Now:      time.Date(2026, 6, 20, 0, 0, 0, 0, time.UTC),
	}
	if err := Run(c, s, opt, &out); err != nil {
		t.Fatalf("Run error: %v", err)
	}
	if len(s.Ingestion.ProcessedWarcs) != 1 {
		t.Errorf("processed %d WARCs, want 1 (MaxWarcs)", len(s.Ingestion.ProcessedWarcs))
	}
}
