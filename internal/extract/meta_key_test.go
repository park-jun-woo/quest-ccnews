//ff:func feature=extract type=helper control=iteration dimension=1
//ff:what metaKey가 property 우선/ name 폴백/property trim/둘다 없음을 소문자 정규화로 올바르게 돌려주는지 테이블로 검증한다.

package extract

import (
	"testing"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func TestMetaKey(t *testing.T) {
	cases := []struct {
		name  string
		attrs []html.Attribute
		want  string
	}{
		{"property preferred", []html.Attribute{{Key: "property", Val: "OG:Title"}, {Key: "name", Val: "twitter"}}, "og:title"},
		{"name fallback", []html.Attribute{{Key: "name", Val: " Author "}}, "author"},
		{"trim property", []html.Attribute{{Key: "property", Val: "  article:author  "}}, "article:author"},
		{"neither", []html.Attribute{{Key: "content", Val: "x"}}, ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			n := &html.Node{Type: html.ElementNode, DataAtom: atom.Meta, Attr: c.attrs}
			if got := metaKey(n); got != c.want {
				t.Fatalf("metaKey = %q, want %q", got, c.want)
			}
		})
	}
}
