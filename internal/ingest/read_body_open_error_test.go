//ff:func feature=ingestion type=helper control=sequence level=error
//ff:what ReadBody가 캐시에 없는 WARC 파일을 열 때 에러를 반환하는지 검증한다.

package ingest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestReadBody_OpenError(t *testing.T) {
	c := NewClient("ua", t.TempDir())
	_, err := c.ReadBody(&session.WARCLoc{File: "missing.warc", Offset: 0})
	if err == nil {
		t.Fatal("want error opening missing file")
	}
}
