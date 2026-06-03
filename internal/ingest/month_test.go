//ff:func feature=ingestion type=helper control=sequence
//ff:what Month.String가 "YYYY/MM" 형식(영점 패딩)으로 포매팅하는지 검증한다.

package ingest

import "testing"

func TestMonthString(t *testing.T) {
	if got := (Month{2026, 6}).String(); got != "2026/06" {
		t.Errorf("String() = %q, want 2026/06", got)
	}
	if got := (Month{2016, 12}).String(); got != "2016/12" {
		t.Errorf("String() = %q, want 2016/12", got)
	}
	if got := (Month{2016, 8}).String(); got != "2016/08" {
		t.Errorf("String() = %q, want 2016/08", got)
	}
}
