//ff:func feature=robots type=helper control=sequence
//ff:what 5xx 응답에서 Fetch가 status=unreachable/crawl-not-allowed이고 룰셋이 비는지 검증한다.

package robots

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchUnreachable5xx(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "boom", http.StatusInternalServerError)
	}))
	defer srv.Close()

	c := newTestClient(t, srv)
	rec, rs, err := c.Fetch("example.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rec.Status != "unreachable" || rec.CrawlAllowed {
		t.Errorf("rec = %+v, want unreachable/not-allowed", rec)
	}
	if rs == nil || len(rs.Groups) != 0 {
		t.Errorf("ruleset should be empty for unreachable, got %+v", rs)
	}
}
