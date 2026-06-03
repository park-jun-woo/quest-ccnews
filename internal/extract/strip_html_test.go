//ff:func feature=extract type=helper control=sequence
//ff:what stripHTML이 script/style/head를 제거하고 공백을 정규화하며, 빈 입력은 빈 문자열, template는 드롭하는지 검증한다.

package extract

import "testing"

func TestStripHTML(t *testing.T) {
	t.Run("removes script/style/head and collapses whitespace", func(t *testing.T) {
		in := []byte(`<html><head><title>Head</title><style>.a{}</style></head>` +
			`<body>  Hello   <script>bad()</script> <noscript>NS</noscript>world  </body></html>`)
		got := stripHTML(in)
		if got != "Hello world" {
			t.Fatalf("stripHTML = %q, want %q", got, "Hello world")
		}
	})
	t.Run("empty input still parses", func(t *testing.T) {
		// html.Parse never errors on byte input; empty yields empty text.
		if got := stripHTML(nil); got != "" {
			t.Fatalf("stripHTML(nil) = %q, want empty", got)
		}
	})
	t.Run("template dropped", func(t *testing.T) {
		got := stripHTML([]byte(`<body>keep<template>drop</template></body>`))
		if got != "keep" {
			t.Fatalf("stripHTML = %q, want %q", got, "keep")
		}
	})
}
