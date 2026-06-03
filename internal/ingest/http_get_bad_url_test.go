//ff:func feature=ingestion type=helper control=sequence
//ff:what httpGet가 http.NewRequest를 실패시키는 잘못된 URL에 대해 에러를 반환하는지 검증한다.

package ingest

import (
	"net/http"
	"testing"
)

func TestHttpGetBadURL(t *testing.T) {
	// control character makes http.NewRequest fail
	if _, err := httpGet(http.DefaultClient, "http://\x7f/bad", "ua"); err == nil {
		t.Fatal("want error on bad URL")
	}
}
