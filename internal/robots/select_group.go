//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what 룰셋에서 우리 product token(parkjunwoo-quest)에 맞는 그룹을 고른다. 콕 집은 그룹 없으면 * 그룹. 순수 함수.

package robots

import "strings"

// SelectGroup picks the group that applies to productToken following RFC 9309
// §2.2.1: the most specific matching user-agent wins; a group naming the token
// (prefix of our product token) beats the "*" group, which is the fallback.
// The first matching group of each kind is kept. Returns nil when neither
// matches.
//
// Per Phase004 we evaluate only the product token; no CCBot impersonation.
func SelectGroup(rs *Ruleset, productToken string) *Group {
	token := strings.ToLower(productToken)
	var star, specific *Group
	for i := range rs.Groups {
		g := &rs.Groups[i]
		star, specific = classifyGroup(g, token, star, specific)
	}
	if specific != nil {
		return specific
	}
	return star
}
