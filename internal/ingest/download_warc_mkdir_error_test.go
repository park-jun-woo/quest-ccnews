//ff:func feature=ingestion type=helper control=sequence
//ff:what DownloadWarc가 캐시 디렉터리 생성(MkdirAll) 실패 시 에러를 반환하는지 검증한다.

package ingest

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestDownloadWarcMkdirError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("x"))
	}))
	defer srv.Close()

	// cacheDir under a regular file → MkdirAll fails
	f := filepath.Join(t.TempDir(), "afile")
	if err := os.WriteFile(f, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	c := clientTo(srv, filepath.Join(f, "sub"))
	if _, err := c.DownloadWarc("d/CC-NEWS-x.warc.gz"); err == nil {
		t.Fatal("want MkdirAll error")
	}
}
