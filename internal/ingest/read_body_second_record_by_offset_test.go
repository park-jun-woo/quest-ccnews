//ff:func feature=ingestion type=helper control=sequence
//ff:what ReadBody가 offset 1로 두 번째 레코드 본문을 골라 반환하는지 검증한다.

package ingest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestReadBody_SecondRecordByOffset(t *testing.T) {
	cacheDir, file := writeWarcHTTP(t, "<html>first</html>", "<html>second</html>")
	c := NewClient("ua", cacheDir)
	got, err := c.ReadBody(&session.WARCLoc{File: file, Offset: 1})
	if err != nil {
		t.Fatalf("ReadBody: %v", err)
	}
	if string(got) != "<html>second</html>" {
		t.Errorf("body = %q, want second", got)
	}
}
