//ff:func feature=ingestion type=helper control=sequence
//ff:what headSize가 HEAD 200 + Content-Length 0(유효한 비음수 크기)에 대해 (0,true)를 돌려주는지 검증한다.

package ingest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHeadSizeZeroLength(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "0")
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	c := clientTo(srv, t.TempDir())
	n, ok := c.headSize(srv.URL)
	if !ok {
		t.Fatal("ok = false, want true")
	}
	if n != 0 {
		t.Errorf("size = %d, want 0", n)
	}
}
