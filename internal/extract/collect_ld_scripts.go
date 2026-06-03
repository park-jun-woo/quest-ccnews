//ff:func feature=extract type=helper control=iteration dimension=1
//ff:what 노드 트리를 순회하며 ld+json script 블록의 비어있지 않은 텍스트를 out에 문서 순서대로 모은다. jsonLDScripts의 재귀 헬퍼.

package extract

import (
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// collectLDScripts walks the node tree, appending the inner text of every
// non-empty <script type="application/ld+json"> block to out in document order.
// Recursive depth-1 helper of jsonLDScripts.
func collectLDScripts(n *html.Node, out *[]string) {
	if n.Type == html.ElementNode && n.DataAtom == atom.Script && isLDJSON(n) {
		if txt := textContent(n); strings.TrimSpace(txt) != "" {
			*out = append(*out, txt)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		collectLDScripts(c, out)
	}
}
