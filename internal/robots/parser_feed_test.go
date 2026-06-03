//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what feed가 field별로 user-agent→그룹, allow/disallow→규칙, crawl-delay→딜레이, unknown→agent run 종료로 디스패치하는지 테이블로 검증한다.

package robots

import "testing"

func TestParserFeed(t *testing.T) {
	// inGroup returns a parser already inside a fresh group (cur set), as the
	// non-user-agent fields require; user-agent opens its own group.
	inGroup := func() *parser {
		p := &parser{rs: &Ruleset{Groups: []Group{{}}}, expectingAgent: true}
		p.cur = &p.rs.Groups[0]
		return p
	}
	fresh := func() *parser { return &parser{rs: &Ruleset{}} }

	tests := []struct {
		name  string
		setup func() *parser
		field string
		value string
		check func(t *testing.T, p *parser)
	}{
		{"user-agent opens group", fresh, "user-agent", "Bot", func(t *testing.T, p *parser) {
			if len(p.rs.Groups) != 1 || p.rs.Groups[0].Agents[0] != "bot" {
				t.Errorf("groups = %+v", p.rs.Groups)
			}
		}},
		{"allow adds allow rule", inGroup, "allow", "/a", func(t *testing.T, p *parser) {
			if len(p.cur.Rules) != 1 || !p.cur.Rules[0].Allow || p.cur.Rules[0].Pattern != "/a" {
				t.Errorf("rules = %+v", p.cur.Rules)
			}
		}},
		{"disallow adds disallow rule", inGroup, "disallow", "/d", func(t *testing.T, p *parser) {
			if len(p.cur.Rules) != 1 || p.cur.Rules[0].Allow {
				t.Errorf("rules = %+v", p.cur.Rules)
			}
		}},
		{"crawl-delay sets delay", inGroup, "crawl-delay", "5", func(t *testing.T, p *parser) {
			if !p.cur.HasDelay || p.cur.CrawlDelay != 5 {
				t.Errorf("delay = %d hasDelay=%v", p.cur.CrawlDelay, p.cur.HasDelay)
			}
		}},
		{"unknown field ends agent run", inGroup, "sitemap", "x", func(t *testing.T, p *parser) {
			if p.expectingAgent {
				t.Errorf("expectingAgent should be false after unknown field")
			}
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.setup()
			p.feed(tt.field, tt.value)
			tt.check(t, p)
		})
	}
}
