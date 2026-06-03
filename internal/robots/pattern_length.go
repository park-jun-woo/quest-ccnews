//ff:func feature=robots type=helper control=sequence
//ff:what 최장 일치 비교용 패턴 길이. 끝 "$" 앵커는 위치 표시이므로 길이에서 제외한다(RFC 9309). 순수 함수.

package robots

// patternLength is the comparison length for longest-match: the pattern's
// character count excluding a trailing "$" anchor (which is a position marker,
// not matchable length per RFC 9309).
func patternLength(pattern string) int {
	if len(pattern) > 0 && pattern[len(pattern)-1] == '$' {
		return len(pattern) - 1
	}
	return len(pattern)
}
