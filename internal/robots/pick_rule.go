//ff:func feature=robots type=helper control=selection
//ff:what 한 규칙이 path에 매칭되면 현재 best와 비교해 더 긴 패턴(동률이면 Allow)을 골라 갱신한다. 순수 함수.

package robots

// pickRule applies RFC 9309 §2.2.2 longest-match precedence for one candidate
// rule against the normalized path: a longer matching pattern wins, and on a
// tie Allow beats Disallow. It returns the (possibly updated) best rule and its
// comparison length. Pure — no IO.
func pickRule(r *Rule, norm string, best *Rule, bestLen int) (*Rule, int) {
	switch {
	case !MatchPattern(r.Pattern, norm):
		return best, bestLen
	case patternLength(r.Pattern) > bestLen:
		return r, patternLength(r.Pattern)
	case patternLength(r.Pattern) == bestLen && r.Allow && (best == nil || !best.Allow):
		return r, bestLen
	default:
		return best, bestLen
	}
}
