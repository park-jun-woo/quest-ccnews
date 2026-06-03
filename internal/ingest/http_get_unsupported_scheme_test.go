//ff:func feature=ingestion type=helper control=sequence
//ff:what httpGet가 지원되지 않는 scheme(ftp) URL에 대해 에러를 반환하는지 검증한다.

package ingest

import (
	"net/http"
	"testing"
)

func TestHttpGetURL(t *testing.T) {
	// a URL that is valid for NewRequest but has an unsupported scheme
	if _, err := httpGet(http.DefaultClient, "ftp://example.invalid/x", "ua"); err == nil {
		t.Fatal("want error for unsupported scheme")
	}
}
