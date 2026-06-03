//ff:func feature=ingestion type=helper control=sequence
//ff:what SelectNext가 시작 월에서 미처리 WARC를 찾아 반환하는지 검증한다.

package ingest

import "testing"

func TestSelectNextFoundCurrentMonth(t *testing.T) {
	start := Month{2026, 6}
	fetch := fakeFetch(map[Month][]string{start: {"d/CC-NEWS-1.warc.gz"}}, nil)
	p, done, err := SelectNext(start, map[string]bool{}, fetch)
	if err != nil || done || p != "d/CC-NEWS-1.warc.gz" {
		t.Errorf("got %q,%v,%v", p, done, err)
	}
}
