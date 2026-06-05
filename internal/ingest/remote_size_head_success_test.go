//ff:func feature=ingestion type=helper control=sequence
//ff:what remoteSize가 Range GET 폴백 없이 HEAD Content-Length를 돌려주는지 검증한다.

package ingest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRemoteSizeHeadSuccess(t *testing.T) {
	const obj = "crawl-data/CC-NEWS/2026/06/CC-NEWS-x.warc.gz"
	const want = int64(9999)
	var rangeProbed bool
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodHead:
			w.Header().Set("Content-Length", fmt.Sprintf("%d", want))
			w.WriteHeader(http.StatusOK)
		case http.MethodGet:
			rangeProbed = true
			w.Header().Set("Content-Range", fmt.Sprintf("bytes 0-0/%d", want))
			w.WriteHeader(http.StatusPartialContent)
			w.Write([]byte("x"))
		}
	}))
	defer srv.Close()

	c := clientTo(srv, t.TempDir())
	n, ok := c.remoteSize(obj)
	if !ok {
		t.Fatal("ok = false, want true")
	}
	if n != want {
		t.Errorf("size = %d, want %d", n, want)
	}
	if rangeProbed {
		t.Error("Range GET was performed; HEAD should have sufficed")
	}
}
