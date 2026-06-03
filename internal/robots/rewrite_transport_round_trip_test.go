//ff:func feature=robots type=helper control=sequence
//ff:what rewriteTransport.RoundTrip: 요청 URL의 scheme/host를 테스트 서버로 바꿔 기본 트랜스포트로 보낸다.

package robots

import "net/http"

func (rt *rewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Scheme = rt.base.Scheme
	req.URL.Host = rt.base.Host
	return http.DefaultTransport.RoundTrip(req)
}
