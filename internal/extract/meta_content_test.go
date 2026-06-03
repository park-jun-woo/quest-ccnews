//ff:func feature=extract type=helper control=sequence
//ff:what metaContent가 content 속성을 trim해 돌려주고, 없으면 빈 문자열인지 검증한다.

package extract

import (
	"testing"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func TestMetaContent(t *testing.T) {
	n := &html.Node{Type: html.ElementNode, DataAtom: atom.Meta, Attr: []html.Attribute{
		{Key: "property", Val: "og:title"},
		{Key: "CONTENT", Val: "  Hello  "},
	}}
	if got := metaContent(n); got != "Hello" {
		t.Fatalf("metaContent = %q, want %q", got, "Hello")
	}
	none := &html.Node{Type: html.ElementNode, DataAtom: atom.Meta, Attr: []html.Attribute{{Key: "property", Val: "x"}}}
	if got := metaContent(none); got != "" {
		t.Fatalf("metaContent none = %q", got)
	}
}
