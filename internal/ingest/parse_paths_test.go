//ff:func feature=ingestion type=helper control=sequence
//ff:what ParsePaths가 개행 구분 본문에서 공백·빈 줄을 제거한 경로 목록을 만드는지 검증한다.

package ingest

import (
	"reflect"
	"testing"
)

func TestParsePaths(t *testing.T) {
	body := "crawl-data/CC-NEWS/2026/06/a.warc.gz\n\n  crawl-data/CC-NEWS/2026/06/b.warc.gz  \n   \n"
	got := ParsePaths(body)
	want := []string{
		"crawl-data/CC-NEWS/2026/06/a.warc.gz",
		"crawl-data/CC-NEWS/2026/06/b.warc.gz",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ParsePaths = %v, want %v", got, want)
	}
}
