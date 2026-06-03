//ff:func feature=ingestion type=helper control=sequence
//ff:what httpGet가 200이 아닌 응답(404)에 대해 에러를 반환하는지 검증한다.

package ingest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpGetNon200(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	if _, err := httpGet(srv.Client(), srv.URL, "ua"); err == nil {
		t.Fatal("want error on 404")
	}
}
