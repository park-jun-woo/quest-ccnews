//ff:func feature=robots type=helper control=sequence
//ff:what 이후 literal 세그먼트를 path[pos:]에서 앞으로 검색해 찾으면 그 끝 위치를 돌려준다. 순수 함수.

package robots

import "strings"

// matchAnywhere searches for seg anywhere in path[pos:] (used for segments after
// the first, which follow a wildcard) and returns the position just past the
// match on success. Pure — no IO.
func matchAnywhere(path string, pos int, seg string) (newPos int, ok bool) {
	idx := strings.Index(path[pos:], seg)
	if idx < 0 {
		return pos, false
	}
	return pos + idx + len(seg), true
}
