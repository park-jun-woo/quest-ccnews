//ff:func feature=robots type=helper control=sequence
//ff:what mergeGroups가 다중 그룹의 Agents·Rules를 선언 순서대로 합집합으로 병합하는지 검증한다.

package robots

import (
	"reflect"
	"testing"
)

func TestMergeGroupsUnionInDeclarationOrder(t *testing.T) {
	a := &Group{
		Agents: []string{"alpha"},
		Rules:  []Rule{{Allow: true, Pattern: "/a"}, {Pattern: "/b"}},
	}
	b := &Group{
		Agents: []string{"beta", "gamma"},
		Rules:  []Rule{{Pattern: "/c"}},
	}
	g := mergeGroups([]*Group{a, b})
	if g == nil {
		t.Fatal("mergeGroups(multi) = nil")
	}
	wantAgents := []string{"alpha", "beta", "gamma"}
	if !reflect.DeepEqual(g.Agents, wantAgents) {
		t.Errorf("Agents = %v, want %v", g.Agents, wantAgents)
	}
	wantRules := []Rule{{Allow: true, Pattern: "/a"}, {Pattern: "/b"}, {Pattern: "/c"}}
	if !reflect.DeepEqual(g.Rules, wantRules) {
		t.Errorf("Rules = %v, want %v", g.Rules, wantRules)
	}
}
