//ff:func feature=robots type=helper control=sequence
//ff:what mergeGroups가 어떤 그룹도 delay를 선언하지 않으면 HasDelay false·CrawlDelay 0을 유지하는지 검증한다.

package robots

import "testing"

func TestMergeGroupsNoDelay(t *testing.T) {
	a := &Group{Agents: []string{"*"}, Rules: []Rule{{Pattern: "/a"}}}
	b := &Group{Agents: []string{"*"}, Rules: []Rule{{Pattern: "/b"}}}
	g := mergeGroups([]*Group{a, b})
	if g.HasDelay {
		t.Errorf("HasDelay = true, want false when no group declares a delay")
	}
	if g.CrawlDelay != 0 {
		t.Errorf("CrawlDelay = %d, want 0", g.CrawlDelay)
	}
}
