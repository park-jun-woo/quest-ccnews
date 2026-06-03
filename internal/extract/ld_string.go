//ff:func feature=extract type=helper control=selection
//ff:what JSON-LD 값을 평문 문자열로 강제변환한다. 문자열은 그대로, 비어있지 않은 배열은 첫 stringable 원소, 그 외는 빈 문자열. 순수 함수.

package extract

import "strings"

// ldString coerces a JSON-LD value to a plain string: a string passes through; a
// non-empty array uses its first stringable element; anything else yields "".
func ldString(v any) string {
	switch t := v.(type) {
	case string:
		return strings.TrimSpace(t)
	case []any:
		if len(t) > 0 {
			return ldString(t[0])
		}
	}
	return ""
}
