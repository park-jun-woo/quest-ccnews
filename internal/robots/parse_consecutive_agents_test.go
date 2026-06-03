//ff:func feature=robots type=helper control=sequence
//ff:what 연속된 user-agent 줄들이 하나의 그룹을 공유하고 소문자로 저장되는지 Parse로 검증한다.

package robots

import "testing"

func TestParseConsecutiveAgentsShareGroup(t *testing.T) {
	content := []byte("User-agent: a\nUser-agent: b\nDisallow: /x\n")
	rs := Parse(content)
	if len(rs.Groups) != 1 {
		t.Fatalf("groups = %d, want 1", len(rs.Groups))
	}
	if len(rs.Groups[0].Agents) != 2 {
		t.Errorf("agents = %v, want 2 (a,b)", rs.Groups[0].Agents)
	}
	if rs.Groups[0].Agents[0] != "a" || rs.Groups[0].Agents[1] != "b" {
		t.Errorf("agents lowercased mismatch: %v", rs.Groups[0].Agents)
	}
}
