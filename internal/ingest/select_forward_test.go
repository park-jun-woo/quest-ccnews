//ff:func feature=ingestion type=helper control=sequence
//ff:what SelectForward가 현재 월에서 미처리 WARC를 찾아 반환하는지 검증한다.

package ingest

import "testing"

func TestSelectForwardFound(t *testing.T) {
	cur := Month{2026, 6}
	fetch := fakeFetch(map[Month][]string{cur: {"d/CC-NEWS-1.warc.gz", "d/CC-NEWS-2.warc.gz"}}, nil)
	p, done, err := SelectForward(cur, map[string]bool{}, fetch)
	if err != nil || done || p != "d/CC-NEWS-2.warc.gz" {
		t.Errorf("got %q,%v,%v", p, done, err)
	}
}
