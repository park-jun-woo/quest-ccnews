//ff:func feature=ingestion type=helper control=sequence
//ff:what 동일 크기 캐시 파일이면 DownloadWarc가 HEAD 탐침으로 재사용하고 본문 GET을 하지 않는지 검증한다.

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

func TestDownloadWarcCacheHitSkipsBody(t *testing.T) {
	payload := []byte("WARC-FAKE-BODY") // 14 bytes
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
	if err := os.WriteFile(dest, payload, 0o644); err != nil {
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
	if n := atomic.LoadInt32(&bodyGets); n != 0 {
		t.Errorf("body GET count = %d, want 0 (cache hit)", n)
	}
	b, _ := os.ReadFile(got)
	if string(b) != string(payload) {
		t.Errorf("cache content mutated = %q", b)
	}
}
