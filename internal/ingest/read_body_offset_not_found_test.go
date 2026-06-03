//ff:func feature=ingestion type=helper control=sequence level=error
//ff:what ReadBody가 레코드 수를 넘는 offset에 대해 "not found" 에러를 반환하는지 검증한다.

package ingest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestReadBody_OffsetNotFound(t *testing.T) {
	cacheDir, file := writeWarcHTTP(t, "<html>only</html>")
	c := NewClient("ua", cacheDir)
	_, err := c.ReadBody(&session.WARCLoc{File: file, Offset: 99})
	if err == nil {
		t.Fatal("want error for offset past end")
	}
	if !strings.Contains(err.Error(), "not found") {
		t.Errorf("err = %v, want 'not found'", err)
	}
}
