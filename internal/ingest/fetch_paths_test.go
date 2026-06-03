//ff:func feature=ingestion type=helper control=sequence
//ff:what FetchPaths가 warc.paths.gz를 받아 gunzip 후 경로 목록으로 파싱하는지 검증한다.

package ingest

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestFetchPathsOK(t *testing.T) {
	body := gzBytes(t, "crawl-data/CC-NEWS/2026/06/a.warc.gz\ncrawl-data/CC-NEWS/2026/06/b.warc.gz\n")
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(body)
	}))
	defer srv.Close()

	c := clientTo(srv, t.TempDir())
	got, err := c.FetchPaths(Month{2026, 6})
	if err != nil {
		t.Fatalf("FetchPaths error: %v", err)
	}
	want := []string{"crawl-data/CC-NEWS/2026/06/a.warc.gz", "crawl-data/CC-NEWS/2026/06/b.warc.gz"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
