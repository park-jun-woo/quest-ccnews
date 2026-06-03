//ff:func feature=ingestion type=helper control=sequence
//ff:what httpGet가 닫힌 서버(연결 거부)에 대해 트랜스포트 에러를 반환하는지 검증한다.

package ingest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpGetTransportError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	srv.Close() // closed → connection refused
	if _, err := httpGet(srv.Client(), srv.URL, "ua"); err == nil {
		t.Fatal("want error on closed server")
	}
}
