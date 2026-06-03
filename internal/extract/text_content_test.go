//ff:func feature=extract type=helper control=sequence
//ff:what textContent가 직속 텍스트 자식만 이어붙이고(주석 무시) 자식 없으면 빈 문자열인지 검증한다.

package extract

import (
	"testing"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func TestTextContent(t *testing.T) {
	n := &html.Node{Type: html.ElementNode, DataAtom: atom.Script}
	n.AppendChild(&html.Node{Type: html.TextNode, Data: "abc"})
	n.AppendChild(&html.Node{Type: html.CommentNode, Data: "ignored"})
	n.AppendChild(&html.Node{Type: html.TextNode, Data: "def"})
	if got := textContent(n); got != "abcdef" {
		t.Fatalf("textContent = %q, want %q", got, "abcdef")
	}
	empty := &html.Node{Type: html.ElementNode, DataAtom: atom.Script}
	if got := textContent(empty); got != "" {
		t.Fatalf("textContent empty = %q", got)
	}
}
