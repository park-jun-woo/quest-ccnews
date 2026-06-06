//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what 룰셋에서 우리 product token(parkjunwoo-quest)에 맞는 그룹들을 골라 병합한 단일 그룹을 돌려준다. 순수 함수.

package robots

import "strings"

// SelectGroup picks the groups that apply to productToken following RFC 9309
// §2.2.1: the most specific matching user-agent wins; groups naming the token
// (prefix of our product token) beat the "*" groups, which are the fallback.
// Records for the same user-agent are merged (§2.2.1), so all matching groups of
// the winning kind are combined into a single returned Group. Returns nil when
// neither kind matches (preserving the default-allow / no-delay contract). Pure
// — no IO; the input Ruleset is never mutated.
//
// Per Phase004 we evaluate only the product token; no CCBot impersonation.
func SelectGroup(rs *Ruleset, productToken string) *Group {
	token := strings.ToLower(productToken)
	var star, specific []*Group
	for i := range rs.Groups {
		g := &rs.Groups[i]
		star, specific = classifyGroup(g, token, star, specific)
	}
	if len(specific) > 0 {
		return mergeGroups(specific)
	}
	if len(star) > 0 {
		return mergeGroups(star)
	}
	return nil
}
