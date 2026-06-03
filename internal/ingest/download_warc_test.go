//ff:func feature=ingestion type=helper control=sequence
//ff:what DownloadWarc가 WARC를 캐시 디렉터리로 받아 올바른 경로·내용으로 저장하는지 검증한다.

package ingest

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestDownloadWarcOK(t *testing.T) {
	payload := []byte("WARC-FAKE-BODY")
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "14")
		w.Write(payload)
	}))
	defer srv.Close()

	dir := t.TempDir()
	c := clientTo(srv, dir)
	obj := "crawl-data/CC-NEWS/2026/06/CC-NEWS-x.warc.gz"
	got, err := c.DownloadWarc(obj)
	if err != nil {
		t.Fatalf("DownloadWarc error: %v", err)
	}
	want := filepath.Join(dir, "CC-NEWS-x.warc.gz")
	if got != want {
		t.Errorf("path = %q want %q", got, want)
	}
	b, _ := os.ReadFile(got)
	if string(b) != string(payload) {
		t.Errorf("file content = %q", b)
	}
}
