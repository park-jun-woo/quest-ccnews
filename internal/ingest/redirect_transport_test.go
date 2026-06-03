//ff:type feature=ingestion type=helper
//ff:what 테스트용 RoundTripper. 모든 요청을 테스트 서버로 재작성해 baseURL 하드코딩 코드도 네트워크 없이 검증한다.

package ingest

import "net/url"

// redirectTransport rewrites every request to the test server, so code that
// hardcodes baseURL (PathsURL/WarcURL) can be exercised without real network.
type redirectTransport struct {
	target *url.URL
}
