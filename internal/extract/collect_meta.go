//ff:func feature=extract type=helper control=iteration dimension=1
//ff:what 노드 트리를 순회하며 <meta> 요소마다 OG/article 폴백 필드를 f에 반영한다. extractMeta의 재귀 헬퍼.

package extract

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// collectMeta walks the node tree, applying each <meta> element's key/content to
// the accumulating Fields. Recursive depth-1 helper of extractMeta.
func collectMeta(n *html.Node, f *Fields) {
	if n.Type == html.ElementNode && n.DataAtom == atom.Meta {
		applyMeta(f, metaKey(n), metaContent(n))
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		collectMeta(c, f)
	}
}
