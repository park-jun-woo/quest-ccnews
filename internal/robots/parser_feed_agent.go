//ff:func feature=robots type=helper control=sequence
//ff:what user-agent 줄을 처리한다. 직전이 user-agent가 아니면 새 그룹을 열고, 소문자 토큰을 현재 그룹에 추가한다.

package robots

import "strings"

// feedAgent handles a user-agent line: it opens a new group unless the previous
// line was also a user-agent (consecutive agents share a group), then appends
// the lowercased token to the current group.
func (p *parser) feedAgent(value string) {
	if p.cur == nil || !p.expectingAgent {
		p.rs.Groups = append(p.rs.Groups, Group{})
		p.cur = &p.rs.Groups[len(p.rs.Groups)-1]
	}
	p.cur.Agents = append(p.cur.Agents, strings.ToLower(value))
	p.expectingAgent = true
}
