//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what consumeSegment이 빈 세그먼트는 통과, index 0은 head 매칭, 이후는 anywhere 매칭으로 위치를 진행하는지 테이블로 검증한다.

package robots

import "testing"

func TestConsumeSegment(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		pos     int
		index   int
		seg     string
		wantPos int
		wantOK  bool
	}{
		{"empty seg consumes nothing", "/abc", 2, 0, "", 2, true},
		{"head match at start", "/abc", 0, 0, "/ab", 3, true},
		{"head mismatch", "/abc", 0, 0, "xyz", 0, false},
		{"head must be at pos", "/abc", 1, 0, "/ab", 1, false},
		{"anywhere later segment", "/aXXbc", 2, 1, "bc", 6, true},
		{"anywhere not found", "/abc", 1, 1, "zz", 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPos, gotOK := consumeSegment(tt.path, tt.pos, tt.index, tt.seg)
			if gotPos != tt.wantPos || gotOK != tt.wantOK {
				t.Errorf("consumeSegment(%q, %d, %d, %q) = (%d, %v), want (%d, %v)",
					tt.path, tt.pos, tt.index, tt.seg, gotPos, gotOK, tt.wantPos, tt.wantOK)
			}
		})
	}
}
