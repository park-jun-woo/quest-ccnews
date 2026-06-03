//ff:func feature=ingestion type=helper control=sequence
//ff:what 트랙의 다음 스캔 시작 월을 정한다. 커서가 가리키는 WARC의 월(있으면), 없으면 now의 월. 순수 함수.

package ingest

import (
	"time"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// StartMonth picks the month a track should begin scanning from: the month of the
// cursor's current WARC if it carries a parseable name, otherwise the month of
// `now`. Forward and backward both resume from the cursor; a fresh cursor starts
// at the newest (current) month. Pure (no IO).
func StartMonth(cur *session.Cursor, now time.Time) Month {
	if cur != nil && cur.Cursor != "" {
		if m, ok := MonthFromWarc(cur.Cursor); ok {
			return m
		}
	}
	return MonthOf(now)
}
