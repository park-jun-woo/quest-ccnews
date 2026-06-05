//ff:func feature=ingestion type=helper control=sequence
//ff:what rangeSize가 206 + 파싱 가능한 Content-Range에 대해 전체 크기와 (size,true)를 돌려주는지 검증한다.

package ingest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRangeSizeOK(t *testing.T) {
	const want = int64(123456)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Range") != "bytes=0-0" {
			t.Errorf("Range = %q, want bytes=0-0", r.Header.Get("Range"))
		}
		w.Header().Set("Content-Range", fmt.Sprintf("bytes 0-0/%d", want))
		w.WriteHeader(http.StatusPartialContent)
		w.Write([]byte("x"))
	}))
	defer srv.Close()

	c := clientTo(srv, t.TempDir())
	n, ok := c.rangeSize(srv.URL)
	if !ok {
		t.Fatal("ok = false, want true")
	}
	if n != want {
		t.Errorf("size = %d, want %d", n, want)
	}
}
