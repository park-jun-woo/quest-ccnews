//ff:func feature=robots type=helper control=sequence
//ff:what Allow/Disallow 줄을 현재 그룹에 규칙으로 추가한다. user-agent 이전에 나온 지시자는 무시한다.

package robots

// feedRule appends an Allow (allow=true) or Disallow rule to the current group.
// A directive that appears before any user-agent line has no group and is
// ignored. Any rule ends the user-agent accumulation run.
func (p *parser) feedRule(allow bool, value string) {
	if p.cur == nil {
		return // directive before any user-agent — not in a group
	}
	p.cur.Rules = append(p.cur.Rules, Rule{Allow: allow, Pattern: value})
	p.expectingAgent = false
}
