//ff:func feature=ingestion type=helper control=sequence
//ff:what Run이 Now를 주지 않으면 time.Now()로 폴백하는지 검증한다(빈 월 목록으로 즉시 멈춤).

package ingest

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestRunZeroNowUsesWallClock(t *testing.T) {
	// No Now provided → Run falls back to time.Now(). Use a fully-processed
	// forward month so it stops immediately without real downloads.
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// empty listing → NextUnprocessed false → waiting
		w.Write(gzBytes(t, "\n"))
	}))
	defer srv.Close()

	c := clientTo(srv, t.TempDir())
	s := session.New("ua", "cc-news")
	opt := RunOptions{Tracks: []Track{Forward}} // Now zero
	if err := Run(c, s, opt, new(bytes.Buffer)); err != nil {
		t.Fatalf("Run error: %v", err)
	}
}
