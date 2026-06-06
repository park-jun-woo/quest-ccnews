//ff:func feature=robots type=helper control=sequence
//ff:what mergeGroups가 둘 이상의 그룹이 delay를 선언했을 때 첫 명시 delay를 우선 채택하는지 검증한다.

package robots

import "testing"

func TestMergeGroupsFirstDelayWins(t *testing.T) {
	a := &Group{Agents: []string{"*"}, HasDelay: true, CrawlDelay: 2}
	b := &Group{Agents: []string{"*"}, HasDelay: true, CrawlDelay: 5}
	g := mergeGroups([]*Group{a, b})
	if !g.HasDelay {
		t.Fatal("HasDelay = false, want true")
	}
	if g.CrawlDelay != 2 {
		t.Errorf("CrawlDelay = %d, want 2 (first explicit delay wins)", g.CrawlDelay)
	}
}
