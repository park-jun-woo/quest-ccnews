//ff:func feature=robots type=helper control=sequence
//ff:what feedRule이 그룹이 있으면 Allow/Disallow 규칙을 추가하고 agent run을 끝내며, 그룹이 없으면(user-agent 이전) 무시하는지 검증한다.

package robots

import "testing"

func TestParserFeedRule(t *testing.T) {
	t.Run("appends allow rule", func(t *testing.T) {
		p := &parser{rs: &Ruleset{Groups: []Group{{}}}, expectingAgent: true}
		p.cur = &p.rs.Groups[0]
		p.feedRule(true, "/a")
		if len(p.cur.Rules) != 1 || !p.cur.Rules[0].Allow || p.cur.Rules[0].Pattern != "/a" {
			t.Fatalf("rules = %+v", p.cur.Rules)
		}
		if p.expectingAgent {
			t.Errorf("expectingAgent should be false after a rule")
		}
	})

	t.Run("appends disallow rule", func(t *testing.T) {
		p := &parser{rs: &Ruleset{Groups: []Group{{}}}}
		p.cur = &p.rs.Groups[0]
		p.feedRule(false, "/d")
		if len(p.cur.Rules) != 1 || p.cur.Rules[0].Allow {
			t.Fatalf("rules = %+v", p.cur.Rules)
		}
	})

	t.Run("no group ignores rule", func(t *testing.T) {
		p := &parser{rs: &Ruleset{}}
		p.feedRule(false, "/d")
		if len(p.rs.Groups) != 0 {
			t.Errorf("groups = %+v, want none", p.rs.Groups)
		}
	})
}
