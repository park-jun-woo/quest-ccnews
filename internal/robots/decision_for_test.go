//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what decisionFor가 nil은 default-allow, Allow 룰은 "Allow: pat", Disallow 룰은 거부+"Disallow: pat"로 변환하는지 테이블로 검증한다.

package robots

import "testing"

func TestDecisionFor(t *testing.T) {
	tests := []struct {
		name        string
		best        *Rule
		wantAllowed bool
		wantRule    string
	}{
		{"nil is default-allow", nil, true, ""},
		{"allow rule", &Rule{Allow: true, Pattern: "/ok"}, true, "Allow: /ok"},
		{"disallow rule", &Rule{Allow: false, Pattern: "/private/*"}, false, "Disallow: /private/*"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := decisionFor(tt.best)
			if d.Allowed != tt.wantAllowed || d.Rule != tt.wantRule {
				t.Errorf("decisionFor(%+v) = %+v, want allowed=%v rule=%q",
					tt.best, d, tt.wantAllowed, tt.wantRule)
			}
		})
	}
}
