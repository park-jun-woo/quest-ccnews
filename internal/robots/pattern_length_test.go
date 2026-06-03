//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what patternLength가 끝 $ 앵커를 길이에서 제외하고 나머지는 그대로 세는지 테이블로 검증한다.

package robots

import "testing"

func TestPatternLength(t *testing.T) {
	tests := []struct {
		pattern string
		want    int
	}{
		{"", 0},
		{"/abc", 4},
		{"/abc$", 4},
		{"$", 0},
		{"/a*b", 4},
	}
	for _, tt := range tests {
		if got := patternLength(tt.pattern); got != tt.want {
			t.Errorf("patternLength(%q) = %d, want %d", tt.pattern, got, tt.want)
		}
	}
}
