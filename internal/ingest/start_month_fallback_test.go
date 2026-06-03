//ff:func feature=ingestion type=helper control=sequence
//ff:what StartMonth가 nil·빈·파싱불가 커서에 대해 now의 월로 폴백하는지 검증한다.

package ingest

import (
	"testing"
	"time"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestStartMonthFallbackToNow(t *testing.T) {
	now := time.Date(2024, 3, 10, 0, 0, 0, 0, time.UTC)
	// nil cursor
	if got := StartMonth(nil, now); got != (Month{2024, 3}) {
		t.Errorf("StartMonth(nil) = %v, want 2024/03", got)
	}
	// empty cursor name
	if got := StartMonth(&session.Cursor{}, now); got != (Month{2024, 3}) {
		t.Errorf("StartMonth(empty) = %v, want 2024/03", got)
	}
	// unparseable cursor name
	if got := StartMonth(&session.Cursor{Cursor: "garbage.warc.gz"}, now); got != (Month{2024, 3}) {
		t.Errorf("StartMonth(garbage) = %v, want 2024/03", got)
	}
}
