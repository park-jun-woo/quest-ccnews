//ff:func feature=extract type=helper control=sequence
//ff:what extractMeta가 og/article 필드를 모아 ok=true, og:title 없으면 ok=false인지 검증한다. (단위: extractMeta)

package extract

import "testing"

func TestExtractMeta(t *testing.T) {
	t.Run("og fields", func(t *testing.T) {
		in := []byte(`<html><head>` +
			`<meta property="og:title" content="OG Title">` +
			`<meta property="og:site_name" content="OG Site">` +
			`<meta property="article:published_time" content="2026-02-03">` +
			`<meta name="article:author" content="Bob">` +
			`</head><body></body></html>`)
		f, ok := extractMeta(in)
		if !ok {
			t.Fatalf("expected ok=true")
		}
		want := Fields{Title: "OG Title", MediaName: "OG Site", PublishedAt: "2026-02-03", Author: "Bob"}
		if f != want {
			t.Fatalf("got %+v want %+v", f, want)
		}
	})
	t.Run("no og:title -> ok false", func(t *testing.T) {
		in := []byte(`<html><head><meta property="og:site_name" content="S"></head><body></body></html>`)
		_, ok := extractMeta(in)
		if ok {
			t.Fatalf("expected ok=false without og:title")
		}
	})
}
