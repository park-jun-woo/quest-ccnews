//ff:func feature=ingestion type=helper control=sequence
//ff:what 테스트 헬퍼. 본문을 CC-NEWS가 WARC response 레코드에 감싸는 형식의 원시 HTTP/1.1 응답으로 만든다.

package ingest

// httpResponse builds a raw HTTP/1.1 response (status line + headers + body) the
// way CC-NEWS wraps an origin response inside a WARC response record.
func httpResponse(body string) string {
	return "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/html\r\n" +
		"Content-Length: " + itoa(len(body)) + "\r\n" +
		"\r\n" +
		body
}
