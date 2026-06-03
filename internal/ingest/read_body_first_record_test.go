//ff:func feature=ingestion type=helper control=sequence
//ff:what ReadBody가 offset 0의 첫 레코드 HTTP 본문(HTML)을 올바르게 반환하는지 검증한다.

package ingest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestReadBody_FirstRecord(t *testing.T) {
	cacheDir, file := writeWarcHTTP(t, "<html>hello</html>")
	c := NewClient("ua", cacheDir)
	got, err := c.ReadBody(&session.WARCLoc{File: file, Offset: 0})
	if err != nil {
		t.Fatalf("ReadBody: %v", err)
	}
	if string(got) != "<html>hello</html>" {
		t.Errorf("body = %q", got)
	}
}
