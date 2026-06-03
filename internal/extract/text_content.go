//ff:func feature=extract type=helper control=iteration dimension=1
//ff:what 노드의 직속 텍스트 자식들을 이어붙인다(파싱된 트리에서 script 본문은 단일 raw-text 자식). 순수 함수.

package extract

import (
	"strings"

	"golang.org/x/net/html"
)

// textContent concatenates the direct text children of a node (script bodies are
// a single raw-text child in the parsed tree).
func textContent(n *html.Node) string {
	var b strings.Builder
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			b.WriteString(c.Data)
		}
	}
	return b.String()
}
