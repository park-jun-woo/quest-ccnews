//ff:func feature=robots type=helper control=sequence
//ff:what 트랜스포트 에러(닫힌 서버)에서 Fetch가 에러 대신 status=unreachable로 매핑하는지 검증한다.

package robots

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchTransportError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	srv.Close() // closed immediately → transport error on Do

	c := newTestClient(t, srv)
	rec, rs, err := c.Fetch("example.com")
	if err != nil {
		t.Fatalf("transport error should be mapped, not returned: %v", err)
	}
	if rec.Status != "unreachable" || rec.CrawlAllowed {
		t.Errorf("rec = %+v, want unreachable on transport error", rec)
	}
	if rs == nil {
		t.Errorf("ruleset should be non-nil empty")
	}
}
