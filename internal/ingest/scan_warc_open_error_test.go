//ff:func feature=ingestion type=helper control=sequence
//ff:what ScanWarc가 존재하지 않는 파일을 열 때 에러를 반환하는지 검증한다.

package ingest

import (
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestScanWarcOpenError(t *testing.T) {
	c := NewClient("ua", t.TempDir())
	err := c.ScanWarc(filepath.Join(t.TempDir(), "nope.warc"), "n", func(*session.Article) {})
	if err == nil {
		t.Fatal("want error opening missing file")
	}
}
