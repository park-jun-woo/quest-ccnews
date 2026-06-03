//ff:func feature=ingestion type=helper control=sequence
//ff:what WarcURL이 WARC 객체 경로의 다운로드 URL을 만드는지 검증한다.

package ingest

import "testing"

func TestWarcURL(t *testing.T) {
	p := "crawl-data/CC-NEWS/2026/06/CC-NEWS-x.warc.gz"
	want := "https://data.commoncrawl.org/" + p
	if got := WarcURL(p); got != want {
		t.Errorf("WarcURL = %q, want %q", got, want)
	}
}
