//ff:func feature=ingestion type=helper control=sequence
//ff:what redirectTransport.RoundTrip: 요청 URL의 scheme/host를 테스트 서버로 바꿔 기본 트랜스포트로 보낸다.

package ingest

import "net/http"

// RoundTrip rewrites the request scheme/host to the test target and delegates.
func (rt redirectTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Scheme = rt.target.Scheme
	req.URL.Host = rt.target.Host
	return http.DefaultTransport.RoundTrip(req)
}
