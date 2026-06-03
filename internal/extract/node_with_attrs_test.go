//ff:func feature=extract type=helper control=sequence
//ff:what 테스트 헬퍼. 주어진 속성들을 가진 <script> 요소 노드를 만든다.

package extract

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// nodeWithAttrs builds a <script> element node carrying the given attributes.
func nodeWithAttrs(attrs ...html.Attribute) *html.Node {
	return &html.Node{Type: html.ElementNode, DataAtom: atom.Script, Data: "script", Attr: attrs}
}
