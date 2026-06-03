//ff:func feature=ingestion type=helper control=sequence
//ff:what MonthOf가 instant를 UTC 기준 Month로 정규화하는지 검증한다.

package ingest

import (
	"testing"
	"time"
)

func TestMonthOf(t *testing.T) {
	// A non-UTC instant should be normalized to UTC.
	loc := time.FixedZone("UTC+9", 9*3600)
	// 2026-06-01 02:00 +09:00 == 2026-05-31 17:00 UTC
	t1 := time.Date(2026, 6, 1, 2, 0, 0, 0, loc)
	if got := MonthOf(t1); got != (Month{2026, 5}) {
		t.Errorf("MonthOf = %v, want 2026/05", got)
	}

	t2 := time.Date(2020, 2, 15, 12, 0, 0, 0, time.UTC)
	if got := MonthOf(t2); got != (Month{2020, 2}) {
		t.Errorf("MonthOf = %v, want 2020/02", got)
	}
}
