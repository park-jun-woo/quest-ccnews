//ff:func feature=ingestion type=helper control=sequence
//ff:what httpResponseBody가 유효한 HTTP 응답에서 본문 바이트만 정확히 추출하는지 검증한다.

package ingest

import (
	"strings"
	"testing"
)

func TestHTTPResponseBody_OK(t *testing.T) {
	got, err := httpResponseBody(strings.NewReader(httpResponse("payload")))
	if err != nil {
		t.Fatalf("httpResponseBody: %v", err)
	}
	if string(got) != "payload" {
		t.Errorf("body = %q, want payload", got)
	}
}
