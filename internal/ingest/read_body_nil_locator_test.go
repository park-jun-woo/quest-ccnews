//ff:func feature=ingestion type=helper control=sequence level=error
//ff:what ReadBody가 nil 로케이터에 대해 에러를 반환하는지 검증한다.

package ingest

import (
	"testing"
)

func TestReadBody_NilLocator(t *testing.T) {
	c := NewClient("ua", t.TempDir())
	if _, err := c.ReadBody(nil); err == nil {
		t.Fatal("want error for nil locator")
	}
}
