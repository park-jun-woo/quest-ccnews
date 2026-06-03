//ff:func feature=extract type=helper control=iteration dimension=1
//ff:what <meta> 요소의 property/name 속성을 소문자로 정규화해 돌려준다(OG는 property, 일부 사이트는 name). 순수 함수.

package extract

import (
	"strings"

	"golang.org/x/net/html"
)

// metaKey returns the lowercased property/name attribute of a <meta> element
// (OG uses "property", some sites use "name").
func metaKey(n *html.Node) string {
	var prop, name string
	for _, a := range n.Attr {
		switch strings.ToLower(a.Key) {
		case "property":
			prop = a.Val
		case "name":
			name = a.Val
		}
	}
	if prop != "" {
		return strings.ToLower(strings.TrimSpace(prop))
	}
	return strings.ToLower(strings.TrimSpace(name))
}
