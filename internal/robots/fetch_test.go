//ff:func feature=robots type=helper control=sequence
//ff:what 2xx 응답에서 Fetch가 status=ok, crawl-allowed, RobotsURL/FetchedAt/crawl-delay와 파싱된 그룹을 채우는지 검증한다.

package robots

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchOK(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("User-Agent"); got != "parkjunwoo-quest/0.1" {
			t.Errorf("UA header = %q", got)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("User-agent: *\nDisallow: /secret\nCrawl-delay: 7\n"))
	}))
	defer srv.Close()

	c := newTestClient(t, srv)
	rec, rs, err := c.Fetch("example.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rec.Status != "ok" || !rec.CrawlAllowed {
		t.Errorf("rec = %+v, want ok/allowed", rec)
	}
	if rec.RobotsURL != "https://example.com/robots.txt" {
		t.Errorf("RobotsURL = %q", rec.RobotsURL)
	}
	if rec.FetchedAt == "" {
		t.Errorf("FetchedAt should be set")
	}
	if rec.CrawlDelaySec != 7 {
		t.Errorf("CrawlDelaySec = %d, want 7", rec.CrawlDelaySec)
	}
	if rs == nil || len(rs.Groups) != 1 {
		t.Fatalf("ruleset = %+v, want 1 group", rs)
	}
}
