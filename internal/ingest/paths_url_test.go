//ff:func feature=ingestion type=helper control=sequence
//ff:what PathsURL이 한 달의 warc.paths.gz 다운로드 URL을 만드는지 검증한다.

package ingest

import "testing"

func TestPathsURL(t *testing.T) {
	want := "https://data.commoncrawl.org/crawl-data/CC-NEWS/2026/06/warc.paths.gz"
	if got := PathsURL(Month{2026, 6}); got != want {
		t.Errorf("PathsURL = %q, want %q", got, want)
	}
}
