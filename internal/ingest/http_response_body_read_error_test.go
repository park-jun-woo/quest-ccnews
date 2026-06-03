//ff:func feature=ingestion type=helper control=sequence level=error
//ff:what httpResponseBody가 Content-Length보다 본문이 짧아 io.ReadAll이 ErrUnexpectedEOF로 실패할 때 그 에러를 반환하는지 검증한다.

package ingest

import (
	"strings"
	"testing"
)

// TestHTTPResponseBody_ReadError builds a well-formed HTTP response header that
// advertises a Content-Length larger than the bytes that actually follow. The
// status/header block parses fine (so http.ReadResponse succeeds), but the
// subsequent io.ReadAll of the body hits an unexpected EOF — covering the
// read-error branch.
func TestHTTPResponseBody_ReadError(t *testing.T) {
	// Content-Length claims 100 bytes; only 5 are present.
	raw := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/html\r\n" +
		"Content-Length: 100\r\n" +
		"\r\n" +
		"short"
	_, err := httpResponseBody(strings.NewReader(raw))
	if err == nil {
		t.Fatal("want io.ReadAll error on truncated body")
	}
	// This is past the parse stage, so it must NOT be the parse-error message.
	if strings.Contains(err.Error(), "parse WARC HTTP response") {
		t.Errorf("expected body read error, got parse error: %v", err)
	}
}
