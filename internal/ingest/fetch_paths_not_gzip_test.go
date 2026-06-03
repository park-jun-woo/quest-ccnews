//ff:func feature=ingestion type=helper control=sequence
//ff:what FetchPaths가 gzip이 아닌 본문에 대해 에러를 반환하는지 검증한다.

package ingest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchPathsNotGzip(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("not gzip data"))
	}))
	defer srv.Close()

	c := clientTo(srv, t.TempDir())
	if _, err := c.FetchPaths(Month{2026, 6}); err == nil {
		t.Fatal("want error on non-gzip body")
	}
}
