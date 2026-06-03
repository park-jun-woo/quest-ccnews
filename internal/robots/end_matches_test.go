//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what endMatches가 앵커 없으면 항상 true, 마지막이 와일드카드(빈 세그먼트)면 true, 아니면 pos==len(path)로 끝 일치를 요구하는지 테이블로 검증한다.

package robots

import "testing"

func TestEndMatches(t *testing.T) {
	tests := []struct {
		name      string
		segments  []string
		path      string
		pos       int
		anchorEnd bool
		want      bool
	}{
		{"no anchor allows remainder", []string{"/a"}, "/abc", 2, false, true},
		{"trailing wildcard empty segment", []string{"/a", ""}, "/abc", 2, true, true},
		{"anchor exact end", []string{"/abc"}, "/abc", 4, true, true},
		{"anchor not at end", []string{"/a"}, "/abc", 2, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := endMatches(tt.segments, tt.path, tt.pos, tt.anchorEnd); got != tt.want {
				t.Errorf("endMatches(%v, %q, %d, %v) = %v, want %v",
					tt.segments, tt.path, tt.pos, tt.anchorEnd, got, tt.want)
			}
		})
	}
}
