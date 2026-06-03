//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what matchAnywhere가 path[pos:]에서 세그먼트를 찾으면 그 끝 위치를, 못 찾으면 (pos,false)를 돌려주는지 테이블로 검증한다.

package robots

import "testing"

func TestMatchAnywhere(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		pos     int
		seg     string
		wantPos int
		wantOK  bool
	}{
		{"found at start", "/abc", 0, "/a", 2, true},
		{"found after offset", "/aXXbc", 2, "bc", 6, true},
		{"not found", "/abc", 1, "zz", 1, false},
		{"empty seg matches at pos", "/abc", 3, "", 3, true},
		{"seg before pos ignored", "/abcabc", 4, "/a", 4, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPos, gotOK := matchAnywhere(tt.path, tt.pos, tt.seg)
			if gotPos != tt.wantPos || gotOK != tt.wantOK {
				t.Errorf("matchAnywhere(%q, %d, %q) = (%d, %v), want (%d, %v)",
					tt.path, tt.pos, tt.seg, gotPos, gotOK, tt.wantPos, tt.wantOK)
			}
		})
	}
}
