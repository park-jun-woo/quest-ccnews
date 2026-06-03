//ff:func feature=extract type=helper control=iteration dimension=1
//ff:what 노드 트리를 순회하며 텍스트 노드를 (공백으로 구분해) 모으고 비콘텐츠 서브트리는 건너뛴다. stripHTML의 재귀 헬퍼.

package extract

import (
	"strings"

	"golang.org/x/net/html"
)

// collectText walks the node tree, appending text nodes (separated by spaces)
// while skipping non-content subtrees. Recursive depth-1 helper of stripHTML.
func collectText(n *html.Node, b *strings.Builder) {
	if n.Type == html.ElementNode && isNonContent(n.DataAtom) {
		return
	}
	if n.Type == html.TextNode {
		b.WriteString(n.Data)
		b.WriteByte(' ')
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		collectText(c, b)
	}
}
