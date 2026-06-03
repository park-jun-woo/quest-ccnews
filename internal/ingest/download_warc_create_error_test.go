//ff:func feature=ingestion type=helper control=sequence
//ff:what DownloadWarc가 대상 경로에 디렉터리가 선점돼 os.Create가 실패할 때 에러를 반환하는지 검증한다.

package ingest

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestDownloadWarcCreateError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("x"))
	}))
	defer srv.Close()

	dir := t.TempDir()
	// Pre-create a directory exactly where the WARC file should be written, so
	// os.Create fails.
	if err := os.MkdirAll(filepath.Join(dir, "CC-NEWS-x.warc.gz"), 0o755); err != nil {
		t.Fatal(err)
	}
	c := clientTo(srv, dir)
	if _, err := c.DownloadWarc("d/CC-NEWS-x.warc.gz"); err == nil {
		t.Fatal("want os.Create error")
	}
}
