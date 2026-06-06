//ff:func feature=robots type=helper control=sequence
//ff:what mergeGroups가 단일 그룹을 새 Group으로 복사하며 Agents·Rules·delay를 그대로 옮기는지 검증한다.

package robots

import (
	"reflect"
	"testing"
)

func TestMergeGroupsSingle(t *testing.T) {
	in := &Group{
		Agents:     []string{"*"},
		Rules:      []Rule{{Allow: false, Pattern: "/x"}},
		CrawlDelay: 3,
		HasDelay:   true,
	}
	g := mergeGroups([]*Group{in})
	if g == nil {
		t.Fatal("mergeGroups(single) = nil, want non-nil")
	}
	if g == in {
		t.Error("mergeGroups returned the input pointer; must be a new Group")
	}
	if !reflect.DeepEqual(g.Agents, []string{"*"}) {
		t.Errorf("Agents = %v, want [*]", g.Agents)
	}
	if !reflect.DeepEqual(g.Rules, []Rule{{Pattern: "/x"}}) {
		t.Errorf("Rules = %v, want [{false /x}]", g.Rules)
	}
	if !g.HasDelay || g.CrawlDelay != 3 {
		t.Errorf("delay = (%v,%d), want (true,3)", g.HasDelay, g.CrawlDelay)
	}
}
