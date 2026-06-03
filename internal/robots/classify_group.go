//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what 한 그룹의 agent 토큰들을 훑어 * 후보 또는 token-prefix 일치(specific) 후보를 첫 등장 기준으로 갱신해 돌려준다. 순수 함수.

package robots

import "strings"

// classifyGroup inspects one group's agent tokens against the lowercased
// product token and updates the running star/specific candidates, keeping the
// first match of each kind. RFC 9309 §2.2.1: a robots user-agent matches when
// it is a prefix of our product token (e.g. site names "parkjunwoo" and we are
// "parkjunwoo-quest"). Pure — no IO.
func classifyGroup(g *Group, token string, star, specific *Group) (newStar, newSpecific *Group) {
	for _, a := range g.Agents {
		if a == "*" && star == nil {
			star = g
		}
		if a != "*" && specific == nil && strings.HasPrefix(token, a) {
			specific = g
		}
	}
	return star, specific
}
