//ff:func feature=extract type=helper control=iteration dimension=1
//ff:what JSON-LD @type 배열의 원소 중 하나라도 허용된 기사 타입 문자열이면 true를 돌려준다. 순수 함수.

package extract

// anyArticleType reports whether any element of a JSON-LD @type array is a
// string naming an accepted article type. Pure.
func anyArticleType(list []any) bool {
	for _, e := range list {
		if s, ok := e.(string); ok && articleTypes[s] {
			return true
		}
	}
	return false
}
