//ff:func feature=ingestion type=helper control=sequence level=error
//ff:what ScanWarc가 헤더의 Content-Length가 정수가 아닌 깨진 WARC 레코드를 만나면 ReadRecord 에러를 그대로 반환하는지 검증한다.

package ingest

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// TestScanWarcRecordError feeds an uncompressed WARC stream whose record header
// carries a non-numeric Content-Length. http header parsing succeeds but the
// reader's strconv.Atoi on content-length fails, so ReadRecord returns an error
// that ScanWarc must propagate — covering the per-record error branch.
func TestScanWarcRecordError(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "rec.warc")
	// Version line, one header field with a bad Content-Length, blank line.
	raw := "WARC/1.0\r\n" +
		"WARC-Type: response\r\n" +
		"Content-Length: not-a-number\r\n" +
		"\r\n"
	if err := os.WriteFile(path, []byte(raw), 0o644); err != nil {
		t.Fatal(err)
	}
	c := NewClient("ua", dir)
	if err := c.ScanWarc(path, "rec.warc", func(*session.Article) {}); err == nil {
		t.Fatal("want ReadRecord error on bad Content-Length")
	}
}
