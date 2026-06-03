//ff:func feature=extract type=helper control=iteration dimension=1
//ff:what <meta> 요소의 content 속성을 trim해 돌려준다. 없으면 빈 문자열. 순수 함수.

package extract

import (
	"strings"

	"golang.org/x/net/html"
)

// metaContent returns the trimmed content attribute of a <meta> element.
func metaContent(n *html.Node) string {
	for _, a := range n.Attr {
		if strings.EqualFold(a.Key, "content") {
			return strings.TrimSpace(a.Val)
		}
	}
	return ""
}
