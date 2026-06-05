//ff:func feature=ingestion type=helper control=sequence
//ff:what HEAD가 사용 가능한 크기를 못 주면 remoteSize가 Range 탐침으로 폴백해 그 total을 돌려주는지 검증한다.

package ingest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRemoteSizeRangeFallback(t *testing.T) {
	const obj = "crawl-data/CC-NEWS/2026/06/CC-NEWS-x.warc.gz"
	const want = int64(54321)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodHead {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Range", fmt.Sprintf("bytes 0-0/%d", want))
		w.WriteHeader(http.StatusPartialContent)
		w.Write([]byte("x"))
	}))
	defer srv.Close()

	c := clientTo(srv, t.TempDir())
	n, ok := c.remoteSize(obj)
	if !ok {
		t.Fatal("ok = false, want true (range fallback)")
	}
	if n != want {
		t.Errorf("size = %d, want %d", n, want)
	}
}
