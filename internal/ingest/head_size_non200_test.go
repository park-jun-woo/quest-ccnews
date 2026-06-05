//ff:func feature=ingestion type=helper control=sequence
//ff:what headSize가 비200 HEAD 응답에 대해 (0,false)를 돌려주는지 검증한다.

package ingest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHeadSizeNon200(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}))
	defer srv.Close()

	c := clientTo(srv, t.TempDir())
	n, ok := c.headSize(srv.URL)
	if ok {
		t.Errorf("ok = true, want false (status %d)", http.StatusMethodNotAllowed)
	}
	if n != 0 {
		t.Errorf("size = %d, want 0", n)
	}
}
