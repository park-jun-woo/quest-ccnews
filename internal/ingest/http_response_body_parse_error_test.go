//ff:func feature=ingestion type=helper control=sequence level=error
//ff:what httpResponseBody가 유효한 HTTP 응답이 아닌 내용에서 "parse WARC HTTP response" 에러를 반환하는지 검증한다.

package ingest

import (
	"strings"
	"testing"
)

func TestHTTPResponseBody_ParseError(t *testing.T) {
	// Content that is not a valid HTTP response → http.ReadResponse fails.
	_, err := httpResponseBody(strings.NewReader("not an http response at all"))
	if err == nil {
		t.Fatal("want parse error")
	}
	if !strings.Contains(err.Error(), "parse WARC HTTP response") {
		t.Errorf("err = %v", err)
	}
}
