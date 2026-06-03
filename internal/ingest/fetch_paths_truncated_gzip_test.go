//ff:func feature=ingestion type=helper control=sequence
//ff:what FetchPaths가 잘린 gzip 스트림(헤더는 유효, 본문 truncated)을 읽다 에러를 반환하는지 검증한다.

package ingest

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFetchPathsTruncatedGzip(t *testing.T) {
	full := gzBytes(t, strings.Repeat("crawl-data/CC-NEWS/2026/06/a.warc.gz\n", 50))
	// Serve a valid gzip header but truncated body → gzip.NewReader succeeds,
	// io.ReadAll on the stream fails (unexpected EOF).
	truncated := full[:len(full)/2]
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(truncated)
	}))
	defer srv.Close()

	c := clientTo(srv, t.TempDir())
	if _, err := c.FetchPaths(Month{2026, 6}); err == nil {
		t.Fatal("want error reading truncated gzip")
	}
}
