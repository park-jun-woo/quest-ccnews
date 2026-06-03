//ff:func feature=ingestion type=helper control=sequence
//ff:what NextUnprocessed가 미처리 WARC를 newest-first로 고르고, 전부 처리됐거나 빈 목록이면 ok=false인지 검증한다.

package ingest

import "testing"

func TestNextUnprocessed(t *testing.T) {
	paths := []string{
		"d/CC-NEWS-1.warc.gz",
		"d/CC-NEWS-2.warc.gz",
		"d/CC-NEWS-3.warc.gz",
	}
	// none processed → newest (last) first
	if p, ok := NextUnprocessed(paths, map[string]bool{}); !ok || p != "d/CC-NEWS-3.warc.gz" {
		t.Errorf("got %q,%v want newest", p, ok)
	}
	// newest processed → next newest
	processed := map[string]bool{"CC-NEWS-3.warc.gz": true}
	if p, ok := NextUnprocessed(paths, processed); !ok || p != "d/CC-NEWS-2.warc.gz" {
		t.Errorf("got %q,%v want CC-NEWS-2", p, ok)
	}
	// all processed → ok=false
	all := map[string]bool{"CC-NEWS-1.warc.gz": true, "CC-NEWS-2.warc.gz": true, "CC-NEWS-3.warc.gz": true}
	if p, ok := NextUnprocessed(paths, all); ok || p != "" {
		t.Errorf("got %q,%v want '',false", p, ok)
	}
	// empty listing
	if _, ok := NextUnprocessed(nil, map[string]bool{}); ok {
		t.Error("empty listing should yield false")
	}
}
