//ff:func feature=cli type=helper control=sequence
//ff:what 테스트 헬퍼. 본문을 최소 HTTP/1.1 200 응답으로 감싼다(CC-NEWS가 WARC에 저장하는 형식).

package cmd

import "strconv"

// httpResp wraps body in a minimal HTTP/1.1 response, the way CC-NEWS stores it.
func httpResp(body string) string {
	return "HTTP/1.1 200 OK\r\nContent-Type: text/html\r\nContent-Length: " +
		strconv.Itoa(len(body)) + "\r\n\r\n" + body
}
