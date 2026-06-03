//ff:func feature=robots type=helper control=sequence
//ff:what 패턴에서 끝 "$" 앵커를 떼어낸다. 앵커가 있으면 제거한 패턴과 true, 없으면 원본과 false. 순수 함수.

package robots

import "strings"

// splitAnchor strips a trailing "$" end-anchor from a pattern, returning the
// remaining pattern and whether an anchor was present. Pure — no IO.
func splitAnchor(pattern string) (rest string, anchorEnd bool) {
	if strings.HasSuffix(pattern, "$") {
		return pattern[:len(pattern)-1], true
	}
	return pattern, false
}
