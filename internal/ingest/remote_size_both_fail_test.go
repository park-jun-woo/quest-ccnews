//ff:func feature=ingestion type=helper control=sequence
//ff:what HEAD·Range 탐침이 둘 다 크기를 못 정하면 remoteSize가 (0,false)를 돌려주는지 검증한다.

package ingest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRemoteSizeBothFail(t *testing.T) {
	const obj = "crawl-data/CC-NEWS/2026/06/CC-NEWS-x.warc.gz"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// HEAD: non-200. GET (Range): non-206. Neither yields a size.
		http.Error(w, "service unavailable", http.StatusServiceUnavailable)
	}))
	defer srv.Close()

	c := clientTo(srv, t.TempDir())
	n, ok := c.remoteSize(obj)
	if ok {
		t.Error("ok = true, want false (both probes failed)")
	}
	if n != 0 {
		t.Errorf("size = %d, want 0", n)
	}
}
