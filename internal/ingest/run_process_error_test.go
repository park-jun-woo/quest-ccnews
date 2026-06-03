//ff:func feature=ingestion type=helper control=sequence
//ff:what Run이 ProcessTrack에서 난 에러를 그대로 전파하는지 검증한다.

package ingest

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestRunProcessError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	c := clientTo(srv, t.TempDir())
	s := session.New("ua", "cc-news")
	opt := RunOptions{
		Tracks: []Track{Forward},
		Now:    time.Date(2026, 6, 20, 0, 0, 0, 0, time.UTC),
	}
	if err := Run(c, s, opt, new(bytes.Buffer)); err == nil {
		t.Fatal("want error propagated from ProcessTrack")
	}
}
