//ff:func feature=robots type=helper control=sequence
//ff:what 첫 literal 세그먼트가 path[pos:]의 맨 앞에 오는지 확인하고, 맞으면 그 길이만큼 진행한 위치를 돌려준다. 순수 함수.

package robots

import "strings"

// matchHead requires seg to sit at the very start of path[pos:] and returns the
// advanced position on success. Pure — no IO.
func matchHead(path string, pos int, seg string) (newPos int, ok bool) {
	if !strings.HasPrefix(path[pos:], seg) {
		return pos, false
	}
	return pos + len(seg), true
}
