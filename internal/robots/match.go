//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what robots 경로 패턴(* 와일드카드, 끝 $ 앵커)이 path에 매칭되는지 RFC 9309 규칙으로 판정한다. 순수 함수.

package robots

import "strings"

// MatchPattern reports whether an RFC 9309 path pattern matches path. Patterns
// match from the start of the path: "*" matches any run of characters and a
// trailing "$" anchors the end. Without "$" a pattern matches if its literal
// segments occur in order starting at the path head (a prefix-style match).
// An empty pattern matches everything. Pure — no IO.
func MatchPattern(pattern, path string) bool {
	pattern, anchorEnd := splitAnchor(pattern)
	if pattern == "" {
		// "" matches all; "$" (now "") requires the path to be empty.
		return !anchorEnd || path == ""
	}

	segments := strings.Split(pattern, "*")
	pos := 0
	for i, seg := range segments {
		next, ok := consumeSegment(path, pos, i, seg)
		if !ok {
			return false
		}
		pos = next
	}
	return endMatches(segments, path, pos, anchorEnd)
}
