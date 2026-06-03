//ff:func feature=ingestion type=helper control=sequence level=error
//ff:what gzip 매직 바이트로 시작하지만 유효 gzip이 아닌 파일에서 warc.NewReader가 실패할 때 ReadBody가 에러를 반환하는지 검증한다.

package ingest

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestReadBody_NewReaderError(t *testing.T) {
	// File starts with the gzip magic bytes (0x1f 0x8b) but is not a valid gzip
	// stream, so warc.NewReader → gzip.NewReader fails (read_body.go:37).
	cacheDir := t.TempDir()
	file := "bad.warc"
	if err := os.WriteFile(filepath.Join(cacheDir, file), []byte{0x1f, 0x8b, 0x00, 0x00}, 0o644); err != nil {
		t.Fatal(err)
	}
	c := NewClient("ua", cacheDir)
	if _, err := c.ReadBody(&session.WARCLoc{File: file, Offset: 0}); err == nil {
		t.Fatal("want NewReader/gzip error")
	}
}
