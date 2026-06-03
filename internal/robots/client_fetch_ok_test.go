//ff:func feature=robots type=helper control=sequence
//ff:what fetchOK가 2xx 본문을 룰셋으로 파싱하고 status=ok/crawl-allowed와 우리 그룹의 crawl-delay를 레코드에 채우는지 검증한다.

package robots

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestFetchOKHelper(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("User-agent: parkjunwoo-quest\nDisallow: /x\nCrawl-delay: 9\n"))
	}))
	defer srv.Close()

	c := newTestClient(t, srv)
	resp, err := c.http.Get(srv.URL + "/robots.txt")
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	defer resp.Body.Close()

	rec := &session.Robots{}
	gotRec, rs, err := c.fetchOK(rec, resp)
	if err != nil {
		t.Fatalf("fetchOK error: %v", err)
	}
	if gotRec.Status != "ok" || !gotRec.CrawlAllowed {
		t.Errorf("rec = %+v, want ok/allowed", gotRec)
	}
	if gotRec.CrawlDelaySec != 9 {
		t.Errorf("CrawlDelaySec = %d, want 9", gotRec.CrawlDelaySec)
	}
	if rs == nil || len(rs.Groups) != 1 {
		t.Fatalf("ruleset = %+v, want 1 group", rs)
	}
}
