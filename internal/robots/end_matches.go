//ff:func feature=robots type=helper control=selection
//ff:what 끝 $ 앵커 처리. 앵커가 없으면 항상 true. 있으면 마지막이 와일드카드거나 소비 위치가 path 끝과 일치해야 한다. 순수 함수.

package robots

// endMatches resolves the trailing "$" anchor after all literal segments are
// consumed. Without an anchor any remaining path is allowed (prefix match).
// With an anchor the path must end exactly where consumption stopped, unless
// the pattern's last token was a "*" (trailing empty segment). Pure — no IO.
func endMatches(segments []string, path string, pos int, anchorEnd bool) bool {
	switch {
	case !anchorEnd:
		return true
	case segments[len(segments)-1] == "":
		return true
	default:
		return pos == len(path)
	}
}
