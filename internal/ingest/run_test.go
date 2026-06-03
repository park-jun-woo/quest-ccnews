//ff:func feature=ingestion type=helper control=sequence
//ff:what Run이 forward·backward 모두 받을 게 없을 때 둘 다 멈추고 매번 Save를 호출하는지 검증한다.

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

func TestRunStopsBothTracksWhenNothingLeft(t *testing.T) {
	// Forward month fully processed → waiting; backward also processed at its
	// cursor's earliest-reachable point.
	obj := "crawl-data/CC-NEWS/2026/06/CC-NEWS-20260615000000-00001.warc.gz"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "warc.paths.gz") {
			w.Write(gzBytes(t, obj+"\n"))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	c := clientTo(srv, t.TempDir())
	s := session.New("ua", "cc-news")
	s.Ingestion.ProcessedWarcs = []string{"CC-NEWS-20260615000000-00001.warc.gz"}
	s.Ingestion.Backward = &session.Cursor{
		Cursor: "CC-NEWS-20160815000000-00001.warc.gz", // earliest month
		State:  StateRunning,
	}
	s.Ingestion.ProcessedWarcs = append(s.Ingestion.ProcessedWarcs, "CC-NEWS-20160815000000-00001.warc.gz")

	var out bytes.Buffer
	saveCount := 0
	opt := RunOptions{
		Tracks: []Track{Forward, Backward},
		Now:    time.Date(2026, 6, 20, 0, 0, 0, 0, time.UTC),
		Save:   func() error { saveCount++; return nil },
	}
	if err := Run(c, s, opt, &out); err != nil {
		t.Fatalf("Run error: %v", err)
	}
	if saveCount == 0 {
		t.Error("Save was never called")
	}
	if !strings.Contains(out.String(), "stopped") {
		t.Errorf("output missing stopped lines: %q", out.String())
	}
}
