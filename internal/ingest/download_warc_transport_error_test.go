//ff:func feature=ingestion type=helper control=sequence
//ff:what DownloadWarc가 닫힌 서버(연결 거부)에 대해 트랜스포트 에러를 반환하는지 검증한다.

package ingest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDownloadWarcTransportError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	srv.Close()
	c := clientTo(srv, t.TempDir())
	if _, err := c.DownloadWarc("d/CC-NEWS-x.warc.gz"); err == nil {
		t.Fatal("want transport error")
	}
}
