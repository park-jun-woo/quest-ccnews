//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what splitAnchor가 끝 "$" 앵커가 있으면 제거한 패턴과 true를, 없으면 원본과 false를 돌려주는지 테이블로 검증한다.

package robots

import "testing"

func TestSplitAnchor(t *testing.T) {
	tests := []struct {
		name       string
		pattern    string
		wantRest   string
		wantAnchor bool
	}{
		{"trailing anchor", "/a$", "/a", true},
		{"no anchor", "/a", "/a", false},
		{"only anchor", "$", "", true},
		{"empty pattern", "", "", false},
		{"mid dollar not stripped", "/a$b", "/a$b", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rest, anchor := splitAnchor(tt.pattern)
			if rest != tt.wantRest || anchor != tt.wantAnchor {
				t.Errorf("splitAnchor(%q) = (%q, %v), want (%q, %v)",
					tt.pattern, rest, anchor, tt.wantRest, tt.wantAnchor)
			}
		})
	}
}
