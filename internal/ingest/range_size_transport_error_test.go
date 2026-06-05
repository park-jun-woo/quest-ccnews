//ff:func feature=ingestion type=helper control=sequence
//ff:what rangeSize가 트랜스포트 수준 실패에 대해 (0,false)를 돌려주는지 검증한다.

package ingest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRangeSizeTransportError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	url := srv.URL
	c := clientTo(srv, t.TempDir())
	srv.Close()
	n, ok := c.rangeSize(url)
	if ok {
		t.Error("ok = true, want false (transport error)")
	}
	if n != 0 {
		t.Errorf("size = %d, want 0", n)
	}
}
