//ff:func feature=robots type=helper control=sequence
//ff:what mergeGroups가 한 그룹만 delay를 선언했을 때 그 HasDelay·CrawlDelay를 병합 결과로 이월하는지 검증한다.

package robots

import "testing"

func TestMergeGroupsCarriesDelayFromOnlyDelayingGroup(t *testing.T) {
	a := &Group{Agents: []string{"*"}, Rules: []Rule{{Pattern: "/a"}}}
	b := &Group{Agents: []string{"*"}, Rules: []Rule{{Pattern: "/b"}}, HasDelay: true, CrawlDelay: 9}
	g := mergeGroups([]*Group{a, b})
	if !g.HasDelay {
		t.Fatalf("HasDelay = false, want true (delay carried from 2nd group)")
	}
	if g.CrawlDelay != 9 {
		t.Errorf("CrawlDelay = %d, want 9", g.CrawlDelay)
	}
}
