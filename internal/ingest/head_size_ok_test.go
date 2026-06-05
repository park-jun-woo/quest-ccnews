//ff:func feature=ingestion type=helper control=sequence
//ff:what headSize가 HEAD 200 + 비음수 Content-Length에 대해 (size,true)를 돌려주는지 검증한다.

package ingest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHeadSizeOK(t *testing.T) {
	const want = int64(4096)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodHead {
			t.Errorf("method = %s, want HEAD", r.Method)
		}
		w.Header().Set("Content-Length", fmt.Sprintf("%d", want))
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	c := clientTo(srv, t.TempDir())
	n, ok := c.headSize(srv.URL)
	if !ok {
		t.Fatal("ok = false, want true")
	}
	if n != want {
		t.Errorf("size = %d, want %d", n, want)
	}
}
