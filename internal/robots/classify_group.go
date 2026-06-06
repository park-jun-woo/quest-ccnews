//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what 한 그룹의 agent 토큰들을 훑어 * 후보 또는 token-prefix 일치(specific) 후보를 누적 슬라이스에 append해 돌려준다. 순수 함수.

package robots

import "strings"

// classifyGroup inspects one group's agent tokens against the lowercased
// product token and appends the group to the running star/specific candidate
// slices. RFC 9309 §2.2.1: groups for the same user-agent must be merged, so
// every matching group of each kind is accumulated (not just the first). A
// robots user-agent matches when it is a prefix of our product token (e.g. site
// names "parkjunwoo" and we are "parkjunwoo-quest"). Pure — no IO.
func classifyGroup(g *Group, token string, star, specific []*Group) (newStar, newSpecific []*Group) {
	var isStar, isSpecific bool
	for _, a := range g.Agents {
		if a == "*" {
			isStar = true
		}
		if a != "*" && strings.HasPrefix(token, a) {
			isSpecific = true
		}
	}
	// Append each group at most once per kind so a group is never merged twice.
	if isStar {
		star = append(star, g)
	}
	if isSpecific {
		specific = append(specific, g)
	}
	return star, specific
}
