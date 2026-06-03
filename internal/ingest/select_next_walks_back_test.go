//ff:func feature=ingestion type=helper control=sequence
//ff:what SelectNext가 현재 월을 다 처리하면 이전 월로 거슬러 가 다음 WARC를 찾는지 검증한다.

package ingest

import "testing"

func TestSelectNextWalksToPreviousMonth(t *testing.T) {
	start := Month{2026, 6}
	fetch := fakeFetch(map[Month][]string{
		start:     {"d/CC-NEWS-jun.warc.gz"},
		{2026, 5}: {"d/CC-NEWS-may.warc.gz"},
	}, nil)
	// June fully processed → should step back to May
	processed := map[string]bool{"CC-NEWS-jun.warc.gz": true}
	p, done, err := SelectNext(start, processed, fetch)
	if err != nil || done || p != "d/CC-NEWS-may.warc.gz" {
		t.Errorf("got %q,%v,%v want may", p, done, err)
	}
}
