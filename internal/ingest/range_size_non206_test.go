//ff:func feature=ingestion type=helper control=sequence
//ff:what rangeSize가 비206 상태(서버가 Range 무시한 200)에 대해 (0,false)를 돌려주는지 검증한다.

package ingest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRangeSizeNon206(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "10")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("0123456789"))
	}))
	defer srv.Close()

	c := clientTo(srv, t.TempDir())
	n, ok := c.rangeSize(srv.URL)
	if ok {
		t.Error("ok = true, want false (status 200, not 206)")
	}
	if n != 0 {
		t.Errorf("size = %d, want 0", n)
	}
}
