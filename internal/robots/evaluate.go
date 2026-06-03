//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what (룰셋, UA, path) → 허용/거부 + 매칭 룰. RFC 9309 최장 일치 우선, 동률이면 Allow 우선. 순수 함수(네트워크 0).

package robots

// Evaluate is the deterministic robots gate core: given a parsed ruleset, our
// product token, and a path, it returns the access Decision. It applies the
// group selected for the token (or "*"), then picks the matching rule by RFC
// 9309 §2.2.2 longest-match precedence — the rule with the longest pattern
// wins, and on equal length Allow beats Disallow. No matching rule means
// default-allow. Pure — no IO.
func Evaluate(rs *Ruleset, productToken, path string) Decision {
	g := SelectGroup(rs, productToken)
	if g == nil {
		return Decision{Allowed: true}
	}
	norm := NormalizePath(path)

	bestLen := -1
	var best *Rule
	for i := range g.Rules {
		r := &g.Rules[i]
		best, bestLen = pickRule(r, norm, best, bestLen)
	}
	return decisionFor(best)
}
