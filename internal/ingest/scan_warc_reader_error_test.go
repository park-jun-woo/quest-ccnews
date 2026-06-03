//ff:func feature=ingestion type=helper control=sequence level=error
//ff:what ScanWarc가 gzip 매직으로 시작하지만 본체가 깨진 파일을 열면 warc.NewReader(gzip 헤더 파싱) 단계에서 에러를 반환하는지 검증한다.

package ingest

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// TestScanWarcReaderError writes a file that opens fine but whose first two
// bytes are the gzip magic number (0x1f 0x8b) followed by garbage. The slyrz
// reader's decompress step then fails inside gzip.NewReader, so warc.NewReader
// returns an error — covering the NewReader error branch in ScanWarc.
func TestScanWarcReaderError(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "bad.warc.gz")
	// gzip magic + invalid deflate stream.
	if err := os.WriteFile(path, []byte{0x1f, 0x8b, 0x00, 0x01, 0x02, 0x03}, 0o644); err != nil {
		t.Fatal(err)
	}
	c := NewClient("ua", dir)
	if err := c.ScanWarc(path, "bad.warc.gz", func(*session.Article) {}); err == nil {
		t.Fatal("want warc.NewReader error on corrupt gzip")
	}
}
