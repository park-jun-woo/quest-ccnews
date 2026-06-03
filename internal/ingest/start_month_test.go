//ff:func feature=ingestion type=helper control=sequence
//ff:what StartMonth가 커서가 가리키는 WARC의 월을 반환하는지 검증한다.

package ingest

import (
	"testing"
	"time"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestStartMonthFromCursor(t *testing.T) {
	cur := &session.Cursor{Cursor: "CC-NEWS-20260615123000-00042.warc.gz"}
	now := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	if got := StartMonth(cur, now); got != (Month{2026, 6}) {
		t.Errorf("StartMonth = %v, want 2026/06", got)
	}
}
