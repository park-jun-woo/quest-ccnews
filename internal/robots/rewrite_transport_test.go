//ff:type feature=robots type=helper
//ff:what 테스트용 RoundTripper. 모든 요청을 테스트 서버로 재작성해 robots.txt fetch 코드를 네트워크 없이 검증한다.

package robots

import "net/url"

// rewriteTransport redirects every request to the test server, ignoring the
// original https://<host>/robots.txt target so we never hit the live network.
type rewriteTransport struct {
	base *url.URL
}
