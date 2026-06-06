//ff:func feature=robots type=helper control=sequence
//ff:what robots 테스트 공용 헬퍼 denyRuleset. 모든 에이전트(*)에 대해 "/" Disallow 한 규칙을 가진 Ruleset을 만들어 거부 분기((false,"robots <rule>"), OutBlock 단락) 검증에 쓴다.
package ccnewsquest

import "github.com/park-jun-woo/quest-ccnews/internal/robots"

func denyRuleset() *robots.Ruleset {
	return &robots.Ruleset{Groups: []robots.Group{{
		Agents: []string{"*"},
		Rules:  []robots.Rule{{Allow: false, Pattern: "/"}},
	}}}
}
