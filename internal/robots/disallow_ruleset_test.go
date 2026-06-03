//ff:func feature=robots type=helper control=sequence
//ff:what 테스트 헬퍼. "*" 그룹에 Disallow: /secret 하나를 가진 룰셋을 만든다.

package robots

func disallowRuleset() *Ruleset {
	return &Ruleset{Groups: []Group{{Agents: []string{"*"}, Rules: []Rule{
		{Allow: false, Pattern: "/secret"},
	}}}}
}
