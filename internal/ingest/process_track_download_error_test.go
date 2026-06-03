//ff:func feature=ingestion type=helper control=sequence
//ff:what ProcessTrack가 WARC 다운로드 실패(403) 시 에러를 반환하는지 검증한다.

package ingest

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestProcessTrackDownloadError(t *testing.T) {
	obj := "crawl-data/CC-NEWS/2026/06/CC-NEWS-20260615000000-00001.warc.gz"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "warc.paths.gz") {
			w.Write(gzBytes(t, obj+"\n"))
			return
		}
		w.WriteHeader(http.StatusForbidden) // WARC download fails
	}))
	defer srv.Close()

	c := clientTo(srv, t.TempDir())
	s := session.New("ua", "cc-news")
	now := time.Date(2026, 6, 20, 0, 0, 0, 0, time.UTC)
	if _, err := c.ProcessTrack(s, Forward, now); err == nil {
		t.Fatal("want download error")
	}
}
