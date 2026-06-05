//ff:func feature=ingestion type=helper control=sequence
//ff:what headSize가 사용 가능한 Content-Length 없는 HEAD 200(chunked)에 대해 (0,false)를 돌려주는지 검증한다.

package ingest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHeadSizeMissingContentLength(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Force chunked transfer so the response carries no Content-Length and
		// the client reports ContentLength = -1.
		w.Header().Set("Transfer-Encoding", "chunked")
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	c := clientTo(srv, t.TempDir())
	n, ok := c.headSize(srv.URL)
	if ok {
		t.Error("ok = true, want false (no Content-Length)")
	}
	if n != 0 {
		t.Errorf("size = %d, want 0", n)
	}
}
