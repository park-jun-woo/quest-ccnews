//ff:func feature=robots type=helper control=sequence
//ff:what 전체 User-Agent 헤더에서 매칭용 product token(첫 "/" 또는 공백 이전 텍스트)을 뽑는다. 순수 함수.

package robots

import "strings"

// ProductToken extracts the matchable product token from a full User-Agent
// header string: the text before the first "/" or whitespace. Pure — no IO.
func ProductToken(userAgent string) string {
	t := userAgent
	if i := strings.IndexByte(t, '/'); i >= 0 {
		t = t[:i]
	}
	if i := strings.IndexAny(t, " \t"); i >= 0 {
		t = t[:i]
	}
	return strings.TrimSpace(t)
}
