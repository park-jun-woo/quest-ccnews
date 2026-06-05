//ff:func feature=ingestion type=helper control=sequence
//ff:what HEAD 미지원 시 크기 탐침이 Range GET으로 폴백해 완전한 캐시 파일을 본문 다운로드 없이 재사용하는지 검증한다.

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

func TestDownloadWarcCacheHitRangeFallback(t *testing.T) {
	payload := []byte("WARC-FAKE-BODY") // 14 bytes
	var fullGets int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodHead {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// GET: if it's the single-byte Range probe, answer 206 with Content-Range.
		if r.Header.Get("Range") != "" {
			w.Header().Set("Content-Range", fmt.Sprintf("bytes 0-0/%d", len(payload)))
			w.WriteHeader(http.StatusPartialContent)
			w.Write(payload[:1])
			return
		}
		atomic.AddInt32(&fullGets, 1)
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(payload)))
		w.Write(payload)
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
	if n := atomic.LoadInt32(&fullGets); n != 0 {
		t.Errorf("full body GET count = %d, want 0 (range-fallback cache hit)", n)
	}
}
