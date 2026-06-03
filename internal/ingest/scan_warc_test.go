//ff:func feature=ingestion type=helper control=sequence
//ff:what ScanWarc가 response 레코드만 TODO 기사로 만들고, request·host없음은 건너뛰며 host·offset을 채우는지 검증한다.

package ingest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestScanWarcEmitsResponses(t *testing.T) {
	path := writeWarc(t, []warcRecord{
		{Type: "response", URI: "https://a.example.com/1"},
		{Type: "request", URI: "https://a.example.com/1"}, // skipped
		{Type: "response", URI: "https://b.example.com/2"},
		{Type: "response", URI: "/no-host"}, // skipped (no host)
	})

	c := NewClient("ua", t.TempDir())
	var got []*session.Article
	err := c.ScanWarc(path, "CC-NEWS-x.warc.gz", func(a *session.Article) {
		got = append(got, a)
	})
	if err != nil {
		t.Fatalf("ScanWarc error: %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("emitted %d articles, want 2", len(got))
	}
	if got[0].Host != "a.example.com" || got[1].Host != "b.example.com" {
		t.Errorf("hosts = %q,%q", got[0].Host, got[1].Host)
	}
	if got[0].WARC.File != "CC-NEWS-x.warc.gz" {
		t.Errorf("warc file = %q", got[0].WARC.File)
	}
	// Offset is a sequential record ordinal: response#0 at 0, response#2 at 2.
	if got[0].WARC.Offset != 0 || got[1].WARC.Offset != 2 {
		t.Errorf("offsets = %d,%d want 0,2", got[0].WARC.Offset, got[1].WARC.Offset)
	}
}
