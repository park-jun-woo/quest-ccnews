//ff:func feature=ingestion type=helper control=sequence
//ff:what MonthFromWarc가 정상 WARC 이름에서 연/월을 파싱하는지 검증한다.

package ingest

import "testing"

func TestMonthFromWarc(t *testing.T) {
	if m, ok := MonthFromWarc("CC-NEWS-20260615123000-00042.warc.gz"); !ok || m != (Month{2026, 6}) {
		t.Errorf("MonthFromWarc = %v,%v want 2026/06,true", m, ok)
	}
	if m, ok := MonthFromWarc("CC-NEWS-20160801000000-00000.warc.gz"); !ok || m != (Month{2016, 8}) {
		t.Errorf("MonthFromWarc = %v,%v want 2016/08,true", m, ok)
	}
}
