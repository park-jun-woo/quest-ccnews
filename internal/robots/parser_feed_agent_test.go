//ff:func feature=robots type=helper control=sequence
//ff:what feedAgent가 새 그룹을 열거나(직전이 user-agent가 아니면) 같은 그룹에 토큰을 누적하며, 토큰을 소문자화하고 expectingAgent를 세우는지 검증한다.

package robots

import "testing"

func TestParserFeedAgent(t *testing.T) {
	t.Run("opens group and lowercases", func(t *testing.T) {
		p := &parser{rs: &Ruleset{}}
		p.feedAgent("GoogleBot")
		if len(p.rs.Groups) != 1 || p.rs.Groups[0].Agents[0] != "googlebot" {
			t.Fatalf("groups = %+v", p.rs.Groups)
		}
		if !p.expectingAgent {
			t.Errorf("expectingAgent should be true")
		}
	})

	t.Run("consecutive agents share group", func(t *testing.T) {
		p := &parser{rs: &Ruleset{}}
		p.feedAgent("a")
		p.feedAgent("b")
		if len(p.rs.Groups) != 1 || len(p.rs.Groups[0].Agents) != 2 {
			t.Fatalf("groups = %+v, want one group with 2 agents", p.rs.Groups)
		}
	})

	t.Run("agent after non-agent opens new group", func(t *testing.T) {
		p := &parser{rs: &Ruleset{}}
		p.feedAgent("a")
		p.expectingAgent = false // simulate an intervening rule line
		p.feedAgent("b")
		if len(p.rs.Groups) != 2 {
			t.Fatalf("groups = %+v, want 2 groups", p.rs.Groups)
		}
	})
}
