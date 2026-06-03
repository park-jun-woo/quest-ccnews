//ff:func feature=extract type=helper control=iteration dimension=1
//ff:what <script> 요소가 ld+json 타입을 선언하는지 판정한다(대소문자 무시, 파라미터 허용). 순수 함수.

package extract

import (
	"strings"

	"golang.org/x/net/html"
)

// isLDJSON reports whether a <script> element declares the ld+json type
// (case-insensitive, parameters tolerated).
func isLDJSON(n *html.Node) bool {
	for _, a := range n.Attr {
		if strings.EqualFold(a.Key, "type") {
			return strings.HasPrefix(strings.ToLower(strings.TrimSpace(a.Val)), "application/ld+json")
		}
	}
	return false
}
