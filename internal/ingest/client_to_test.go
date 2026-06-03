//ff:func feature=ingestion type=helper control=sequence
//ff:what 테스트 헬퍼. HTTP 트래픽을 테스트 서버로 우회시키는 Client를 만든다.

package ingest

import (
	"net/http/httptest"
	"net/url"
)

// clientTo builds a Client whose HTTP traffic is redirected to srv.
func clientTo(srv *httptest.Server, cacheDir string) *Client {
	u, _ := url.Parse(srv.URL)
	c := NewClient("test-ua/1.0", cacheDir)
	c.http.Transport = redirectTransport{target: u}
	return c
}
