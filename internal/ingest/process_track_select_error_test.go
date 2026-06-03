//ff:func feature=ingestion type=helper control=sequence
//ff:what ProcessTrack가 warc.paths fetch 실패 시 에러를 반환하는지 검증한다.

package ingest

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestProcessTrackSelectError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	c := clientTo(srv, t.TempDir())
	s := session.New("ua", "cc-news")
	now := time.Date(2026, 6, 20, 0, 0, 0, 0, time.UTC)
	if _, err := c.ProcessTrack(s, Forward, now); err == nil {
		t.Fatal("want error when fetch fails")
	}
}
