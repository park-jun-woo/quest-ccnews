//ff:func feature=robots type=helper control=sequence
//ff:what crawl-delay가 없는 2xx 응답에서 Fetch가 status=ok이고 CrawlDelaySec=0인지 검증한다.

package robots

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchOKNoDelay(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("User-agent: *\nDisallow: /secret\n"))
	}))
	defer srv.Close()

	c := newTestClient(t, srv)
	rec, _, err := c.Fetch("example.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rec.Status != "ok" {
		t.Errorf("status = %q, want ok", rec.Status)
	}
	if rec.CrawlDelaySec != 0 {
		t.Errorf("CrawlDelaySec = %d, want 0", rec.CrawlDelaySec)
	}
}
