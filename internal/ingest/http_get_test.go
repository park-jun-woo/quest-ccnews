//ff:func feature=ingestion type=helper control=sequence
//ff:what httpGet가 200 응답에서 user-agent를 붙여 본문 reader를 반환하는지 검증한다.

package ingest

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpGetOK(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("User-Agent") != "ua/9" {
			t.Errorf("missing UA, got %q", r.Header.Get("User-Agent"))
		}
		io.WriteString(w, "hello")
	}))
	defer srv.Close()

	body, err := httpGet(srv.Client(), srv.URL, "ua/9")
	if err != nil {
		t.Fatalf("httpGet error: %v", err)
	}
	defer body.Close()
	b, _ := io.ReadAll(body)
	if string(b) != "hello" {
		t.Errorf("body = %q", b)
	}
}
