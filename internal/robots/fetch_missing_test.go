//ff:func feature=robots type=helper control=sequence
//ff:what 4xx 응답에서 Fetch가 status=missing/crawl-allowed이고 룰셋이 비는지 검증한다.

package robots

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchMissing4xx(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "not found", http.StatusNotFound)
	}))
	defer srv.Close()

	c := newTestClient(t, srv)
	rec, rs, err := c.Fetch("example.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rec.Status != "missing" || !rec.CrawlAllowed {
		t.Errorf("rec = %+v, want missing/allowed", rec)
	}
	if rs == nil || len(rs.Groups) != 0 {
		t.Errorf("ruleset should be empty for missing, got %+v", rs)
	}
}
