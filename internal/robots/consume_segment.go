//ff:func feature=robots type=helper control=selection
//ff:what 와일드카드로 분리된 literal 세그먼트 하나를 path에서 소비한다. 첫 세그먼트는 head에, 이후는 순서대로 등장해야 한다. 순수 함수.

package robots

// consumeSegment advances the match position past one literal segment of a
// wildcard-split pattern. The empty segment (adjacent "*") consumes nothing.
// The first segment (index 0) must sit at the very head of the remaining path;
// later segments may appear anywhere after the current position, in order. It
// returns the new position and whether the segment matched. Pure — no IO.
func consumeSegment(path string, pos, index int, seg string) (newPos int, ok bool) {
	switch {
	case seg == "":
		return pos, true
	case index == 0:
		return matchHead(path, pos, seg)
	default:
		return matchAnywhere(path, pos, seg)
	}
}
