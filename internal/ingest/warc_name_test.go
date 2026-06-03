//ff:func feature=ingestion type=helper control=sequence
//ff:what WarcName이 객체 경로에서 basename을 뽑고, 슬래시가 없으면 그대로 반환하는지 검증한다.

package ingest

import "testing"

func TestWarcName(t *testing.T) {
	if got := WarcName("crawl-data/CC-NEWS/2026/06/CC-NEWS-x.warc.gz"); got != "CC-NEWS-x.warc.gz" {
		t.Errorf("WarcName = %q", got)
	}
	// no slash → returned as-is
	if got := WarcName("CC-NEWS-x.warc.gz"); got != "CC-NEWS-x.warc.gz" {
		t.Errorf("WarcName(no slash) = %q", got)
	}
}
