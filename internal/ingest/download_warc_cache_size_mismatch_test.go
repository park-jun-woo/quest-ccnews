//ff:func feature=ingestion type=helper control=sequence
//ff:what 캐시 파일 크기가 원격 Content-Length와 다르면 DownloadWarc가 새 다운로드로 덮어쓰는지 검증한다.

package ingest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync/atomic"
	"testing"
)

func TestDownloadWarcCacheSizeMismatchRedownloads(t *testing.T) {
	payload := []byte("WARC-FAKE-BODY") // 14 bytes (the correct size)
	var bodyGets int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(payload)))
		switch r.Method {
		case http.MethodHead:
			w.WriteHeader(http.StatusOK)
		case http.MethodGet:
			atomic.AddInt32(&bodyGets, 1)
			w.Write(payload)
		}
	}))
	defer srv.Close()

	dir := t.TempDir()
	obj := "crawl-data/CC-NEWS/2026/06/CC-NEWS-x.warc.gz"
	dest := filepath.Join(dir, "CC-NEWS-x.warc.gz")
	// Seed a truncated/corrupt cache file (wrong size).
	if err := os.WriteFile(dest, []byte("short"), 0o644); err != nil {
		t.Fatalf("seed cache: %v", err)
	}

	c := clientTo(srv, dir)
	got, err := c.DownloadWarc(obj)
	if err != nil {
		t.Fatalf("DownloadWarc error: %v", err)
	}
	if got != dest {
		t.Errorf("path = %q want %q", got, dest)
	}
	if n := atomic.LoadInt32(&bodyGets); n != 1 {
		t.Errorf("body GET count = %d, want 1 (re-download)", n)
	}
	b, _ := os.ReadFile(got)
	if string(b) != string(payload) {
		t.Errorf("file not overwritten with correct content = %q", b)
	}
}
