//ff:func feature=ingestion type=helper control=sequence
//ff:what Month.Prev가 한 달 이전을 반환하고, 1월은 전년 12월로 돌며, 최古(2016-08)에서 ok=false인지 검증한다.

package ingest

import "testing"

func TestMonthPrev(t *testing.T) {
	// mid-year
	if m, ok := (Month{2026, 6}).Prev(); !ok || m != (Month{2026, 5}) {
		t.Errorf("Prev = %v,%v want 2026/05,true", m, ok)
	}
	// January rolls to previous December
	if m, ok := (Month{2026, 1}).Prev(); !ok || m != (Month{2025, 12}) {
		t.Errorf("Prev = %v,%v want 2025/12,true", m, ok)
	}
	// earliest dump: no earlier month
	if m, ok := (Month{firstYear, firstMonth}).Prev(); ok || m != (Month{firstYear, firstMonth}) {
		t.Errorf("Prev at earliest = %v,%v want self,false", m, ok)
	}
}
