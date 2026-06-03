//ff:func feature=robots type=helper control=sequence
//ff:what 테스트 헬퍼. HTTP 요청이 주어진 테스트 서버로 라우팅되는 Client를 만든다.

package robots

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// newTestClient builds a Client whose HTTP requests are routed to srv.
func newTestClient(t *testing.T, srv *httptest.Server) *Client {
	t.Helper()
	u, err := url.Parse(srv.URL)
	if err != nil {
		t.Fatalf("parse server url: %v", err)
	}
	return &Client{
		http:         &http.Client{Transport: &rewriteTransport{base: u}},
		userAgent:    "parkjunwoo-quest/0.1",
		productToken: "parkjunwoo-quest",
	}
}
