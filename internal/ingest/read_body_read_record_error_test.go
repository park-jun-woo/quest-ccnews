//ff:func feature=ingestion type=helper control=sequence level=error
//ff:what 비숫자 Content-Length로 ReadRecord가 비-EOF 에러를 낼 때 ReadBody가 에러를 반환하는지 검증한다.

package ingest

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestReadBody_ReadRecordError(t *testing.T) {
	// A WARC record with a non-numeric Content-Length makes ReadRecord return a
	// non-EOF error (read_body.go:48), exercised before the offset is reached.
	rec := "WARC/1.0\r\n" +
		"WARC-Type: response\r\n" +
		"Content-Length: not-a-number\r\n" +
		"\r\n" +
		"body\r\n\r\n"
	cacheDir := t.TempDir()
	file := "malformed.warc"
	if err := os.WriteFile(filepath.Join(cacheDir, file), []byte(rec), 0o644); err != nil {
		t.Fatal(err)
	}
	c := NewClient("ua", cacheDir)
	if _, err := c.ReadBody(&session.WARCLoc{File: file, Offset: 5}); err == nil {
		t.Fatal("want ReadRecord error from bad Content-Length")
	}
}
