//ff:func feature=ingestion type=helper control=sequence
//ff:what rangeSize가 Content-Range가 음수 total을 광고하는 206을 거부하고 (0,false)를 돌려주는지 검증한다.

package ingest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRangeSizeNegativeTotal(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Range", "bytes 0-0/-5")
		w.WriteHeader(http.StatusPartialContent)
		w.Write([]byte("x"))
	}))
	defer srv.Close()

	c := clientTo(srv, t.TempDir())
	n, ok := c.rangeSize(srv.URL)
	if ok {
		t.Error("ok = true, want false (negative total)")
	}
	if n != 0 {
		t.Errorf("size = %d, want 0", n)
	}
}
