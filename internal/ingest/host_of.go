//ff:func feature=ingestion type=helper control=sequence
//ff:what URL에서 host(소문자, 포트 제외)를 뽑는다. 파싱 실패나 host 부재면 ok=false. 순수 함수.

package ingest

import (
	"net/url"
	"strings"
)

// HostOf extracts the lowercase hostname (without port) from a URL. It returns
// ok=false when the URL cannot be parsed or carries no host. Pure (no IO).
func HostOf(rawURL string) (host string, ok bool) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", false
	}
	h := u.Hostname()
	if h == "" {
		return "", false
	}
	return strings.ToLower(h), true
}
