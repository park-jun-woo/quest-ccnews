//ff:func feature=robots type=helper control=sequence
//ff:what Parse가 단일 * 그룹의 Disallow/Allow 규칙과 crawl-delay를 올바르게 모으는지 검증한다.

package robots

import "testing"

func TestParseBasicGroup(t *testing.T) {
	content := []byte("User-agent: *\nDisallow: /private\nAllow: /private/public\nCrawl-delay: 5\n")
	rs := Parse(content)
	if len(rs.Groups) != 1 {
		t.Fatalf("groups = %d, want 1", len(rs.Groups))
	}
	g := rs.Groups[0]
	if len(g.Agents) != 1 || g.Agents[0] != "*" {
		t.Errorf("agents = %v, want [*]", g.Agents)
	}
	if len(g.Rules) != 2 {
		t.Fatalf("rules = %d, want 2", len(g.Rules))
	}
	if g.Rules[0].Allow || g.Rules[0].Pattern != "/private" {
		t.Errorf("rule0 = %+v", g.Rules[0])
	}
	if !g.Rules[1].Allow || g.Rules[1].Pattern != "/private/public" {
		t.Errorf("rule1 = %+v", g.Rules[1])
	}
	if !g.HasDelay || g.CrawlDelay != 5 {
		t.Errorf("delay HasDelay=%v val=%d, want true/5", g.HasDelay, g.CrawlDelay)
	}
}
