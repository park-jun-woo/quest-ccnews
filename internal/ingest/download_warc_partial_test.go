//ff:func feature=ingestion type=helper control=sequence
//ff:what DownloadWarc가 Content-Length와 받은 바이트가 다르면 부분 다운로드로 거부하고 파일을 지우는지 검증한다.

package ingest

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestDownloadWarcPartial(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// advertise more than we send
		w.Header().Set("Content-Length", "100")
		w.Write([]byte("short"))
	}))
	defer srv.Close()

	dir := t.TempDir()
	c := clientTo(srv, dir)
	if _, err := c.DownloadWarc("d/CC-NEWS-x.warc.gz"); err == nil {
		t.Fatal("want partial-download error")
	}
	// the partial file must have been removed
	if _, err := os.Stat(filepath.Join(dir, "CC-NEWS-x.warc.gz")); !os.IsNotExist(err) {
		t.Error("partial file should be removed")
	}
}
