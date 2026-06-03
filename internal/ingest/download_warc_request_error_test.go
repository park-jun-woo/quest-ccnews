//ff:func feature=ingestion type=helper control=sequence
//ff:what DownloadWarc가 objectPath에 제어문자가 섞여 WarcURL이 파싱 불가한 URL이 되면 http.NewRequest 단계에서 에러를 반환하는지 검증한다.

package ingest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestDownloadWarcRequestError feeds an objectPath containing a control
// character so WarcURL yields an unparseable URL and http.NewRequest fails
// before any network call — covering the request-build error branch.
func TestDownloadWarcRequestError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("x"))
	}))
	defer srv.Close()

	c := clientTo(srv, t.TempDir())
	// A DEL control char in the path makes url.Parse (inside http.NewRequest) fail.
	if _, err := c.DownloadWarc("d/CC-NEWS\x7f.warc.gz"); err == nil {
		t.Fatal("want http.NewRequest error for control-char URL")
	}
}
