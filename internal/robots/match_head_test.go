//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what matchHead가 세그먼트가 path[pos:] 맨 앞에 올 때만 진행 위치를, 아니면 (pos,false)를 돌려주는지 테이블로 검증한다.

package robots

import "testing"

func TestMatchHead(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		pos     int
		seg     string
		wantPos int
		wantOK  bool
	}{
		{"prefix at start", "/abc", 0, "/ab", 3, true},
		{"prefix at offset", "/abc", 1, "ab", 3, true},
		{"not a prefix", "/abc", 0, "ab", 0, false},
		{"empty seg always matches", "/abc", 2, "", 2, true},
		{"exact remainder", "/abc", 0, "/abc", 4, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPos, gotOK := matchHead(tt.path, tt.pos, tt.seg)
			if gotPos != tt.wantPos || gotOK != tt.wantOK {
				t.Errorf("matchHead(%q, %d, %q) = (%d, %v), want (%d, %v)",
					tt.path, tt.pos, tt.seg, gotPos, gotOK, tt.wantPos, tt.wantOK)
			}
		})
	}
}
