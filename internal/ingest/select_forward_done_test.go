//ff:func feature=ingestion type=helper control=sequence
//ff:what SelectForward가 현재 월을 전부 처리했으면 done=true(→waiting)를 반환하는지 검증한다.

package ingest

import "testing"

func TestSelectForwardDone(t *testing.T) {
	cur := Month{2026, 6}
	fetch := fakeFetch(map[Month][]string{cur: {"d/CC-NEWS-1.warc.gz"}}, nil)
	processed := map[string]bool{"CC-NEWS-1.warc.gz": true}
	p, done, err := SelectForward(cur, processed, fetch)
	if err != nil || !done || p != "" {
		t.Errorf("got %q,%v,%v want done", p, done, err)
	}
}
