//ff:func feature=ingestion type=helper control=sequence
//ff:what SelectNext가 최古 월(2016-08)까지 다 처리하면 done=true(→exhausted)를 반환하는지 검증한다.

package ingest

import "testing"

func TestSelectNextExhausted(t *testing.T) {
	start := Month{firstYear, firstMonth} // earliest dump
	fetch := fakeFetch(map[Month][]string{start: {"d/CC-NEWS-x.warc.gz"}}, nil)
	processed := map[string]bool{"CC-NEWS-x.warc.gz": true}
	p, done, err := SelectNext(start, processed, fetch)
	if err != nil || !done || p != "" {
		t.Errorf("got %q,%v,%v want done", p, done, err)
	}
}
