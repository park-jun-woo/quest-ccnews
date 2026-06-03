//ff:func feature=ingestion type=helper control=sequence
//ff:what DownloadWarc가 200이 아닌 응답(403)에 대해 에러를 반환하는지 검증한다.

package ingest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDownloadWarcStatusError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	}))
	defer srv.Close()

	c := clientTo(srv, t.TempDir())
	if _, err := c.DownloadWarc("d/CC-NEWS-x.warc.gz"); err == nil {
		t.Fatal("want error on 403")
	}
}
