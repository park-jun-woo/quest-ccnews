//ff:func feature=extract type=helper control=selection
//ff:what JSON-LD 사람/조직 참조를 표시 이름 문자열로 정규화한다. 문자열/객체(.name)/배열(첫 원소)을 받는다. author·publisher에 사용. 순수 함수.

package extract

import "strings"

// ldName resolves a JSON-LD person/organization reference to a display name. It
// accepts a bare string, an object with a "name" property, or an array (first
// element). Used for author and publisher. Pure.
func ldName(v any) string {
	switch t := v.(type) {
	case string:
		return strings.TrimSpace(t)
	case map[string]any:
		return ldString(t["name"])
	case []any:
		if len(t) > 0 {
			return ldName(t[0])
		}
	}
	return ""
}
