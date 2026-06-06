//ff:func feature=robots type=helper control=sequence
//ff:what Allow/Disallow 줄을 현재 그룹에 규칙으로 추가한다. user-agent 이전 지시자와 빈 값(효력 없음) 지시자는 무시한다.

package robots

import "strings"

// feedRule appends an Allow (allow=true) or Disallow rule to the current group.
// A directive that appears before any user-agent line has no group and is
// ignored. Any rule ends the user-agent accumulation run. Per RFC 9309 an empty
// value (e.g. "Disallow:") has no effect, so it is not turned into a rule —
// adding it as an empty pattern would otherwise match every path. The agent
// accumulation still ends. This is symmetric for Allow and Disallow.
func (p *parser) feedRule(allow bool, value string) {
	if p.cur == nil {
		return // directive before any user-agent — not in a group
	}
	if strings.TrimSpace(value) == "" {
		p.expectingAgent = false // end UA run, but an empty value adds no rule
		return
	}
	p.cur.Rules = append(p.cur.Rules, Rule{Allow: allow, Pattern: value})
	p.expectingAgent = false
}
