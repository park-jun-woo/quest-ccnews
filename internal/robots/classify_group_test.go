//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what classifyGroup이 * 후보와 token-prefix(specific) 후보를 누적 슬라이스에 append하는지 테이블로 검증한다.

package robots

import "testing"

func TestClassifyGroup(t *testing.T) {
	starA := &Group{Agents: []string{"*"}}
	specA := &Group{Agents: []string{"parkjunwoo"}}
	specB := &Group{Agents: []string{"parkjunwoo-quest"}}
	mixed := &Group{Agents: []string{"*", "parkjunwoo"}}
	other := &Group{Agents: []string{"googlebot"}}

	tests := []struct {
		name         string
		g            *Group
		token        string
		star, spec   []*Group
		wantStar     []*Group
		wantSpecific []*Group
	}{
		{"star appends", starA, "parkjunwoo-quest", nil, nil, []*Group{starA}, nil},
		{"star accumulates", starA, "parkjunwoo-quest", []*Group{starA}, nil, []*Group{starA, starA}, nil},
		{"specific prefix match", specA, "parkjunwoo-quest", nil, nil, nil, []*Group{specA}},
		{"exact token match", specB, "parkjunwoo-quest", nil, nil, nil, []*Group{specB}},
		{"specific accumulates", specB, "parkjunwoo-quest", nil, []*Group{specA}, nil, []*Group{specA, specB}},
		{"no match leaves slices", other, "parkjunwoo-quest", nil, nil, nil, nil},
		{"mixed appends both", mixed, "parkjunwoo-quest", nil, nil, []*Group{mixed}, []*Group{mixed}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStar, gotSpec := classifyGroup(tt.g, tt.token, tt.star, tt.spec)
			gs, gp := ptrs(gotStar), ptrs(gotSpec)
			if gs != ptrs(tt.wantStar) || gp != ptrs(tt.wantSpecific) {
				t.Errorf("classifyGroup() = (star=%s, spec=%s), want (star=%s, spec=%s)",
					gs, gp, ptrs(tt.wantStar), ptrs(tt.wantSpecific))
			}
		})
	}
}
