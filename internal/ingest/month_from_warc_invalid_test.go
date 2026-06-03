//ff:func feature=ingestion type=helper control=iteration dimension=1
//ff:what MonthFromWarc가 접두사 불일치·길이 부족·비수치·월 범위 밖 이름에 대해 ok=false를 반환하는지 검증한다.

package ingest

import "testing"

func TestMonthFromWarcInvalid(t *testing.T) {
	bad := []string{
		"OTHER-20260615123000.warc.gz", // wrong prefix
		"CC-NEWS-2026",                 // too short
		"CC-NEWS-XXXX0615.warc.gz",     // year not numeric
		"CC-NEWS-2026XX15.warc.gz",     // month not numeric
		"CC-NEWS-20261315.warc.gz",     // month out of range (13)
		"CC-NEWS-20260015.warc.gz",     // month 0
	}
	for _, name := range bad {
		if m, ok := MonthFromWarc(name); ok || m != (Month{}) {
			t.Errorf("MonthFromWarc(%q) = %v,%v want zero,false", name, m, ok)
		}
	}
}
