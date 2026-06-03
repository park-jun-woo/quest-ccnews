//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what classifyGroup이 * 후보와 token-prefix(specific) 후보를 첫 등장 기준으로만 갱신하는지 테이블로 검증한다.

package robots

import "testing"

func TestClassifyGroup(t *testing.T) {
	starA := &Group{Agents: []string{"*"}}
	starB := &Group{Agents: []string{"*"}}
	specA := &Group{Agents: []string{"parkjunwoo"}}
	specB := &Group{Agents: []string{"parkjunwoo-quest"}}
	mixed := &Group{Agents: []string{"*", "parkjunwoo"}}
	other := &Group{Agents: []string{"googlebot"}}

	tests := []struct {
		name         string
		g            *Group
		token        string
		star, spec   *Group
		wantStar     *Group
		wantSpecific *Group
	}{
		{"star sets star", starA, "parkjunwoo-quest", nil, nil, starA, nil},
		{"star kept first", starB, "parkjunwoo-quest", starA, nil, starA, nil},
		{"specific prefix match", specA, "parkjunwoo-quest", nil, nil, nil, specA},
		{"exact token match", specB, "parkjunwoo-quest", nil, nil, nil, specB},
		{"specific kept first", specB, "parkjunwoo-quest", nil, specA, nil, specA},
		{"no match leaves nil", other, "parkjunwoo-quest", nil, nil, nil, nil},
		{"mixed sets both", mixed, "parkjunwoo-quest", nil, nil, mixed, mixed},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStar, gotSpec := classifyGroup(tt.g, tt.token, tt.star, tt.spec)
			if gotStar != tt.wantStar || gotSpec != tt.wantSpecific {
				t.Errorf("classifyGroup() = (star=%p, spec=%p), want (star=%p, spec=%p)",
					gotStar, gotSpec, tt.wantStar, tt.wantSpecific)
			}
		})
	}
}
