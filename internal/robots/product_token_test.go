//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what ProductToken이 "/"·공백·탭 이전 토큰을 뽑고 선행 공백/빈 입력을 처리하는지 테이블로 검증한다.

package robots

import "testing"

func TestProductToken(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"parkjunwoo-quest/0.1 (+url)", "parkjunwoo-quest"},
		{"parkjunwoo-quest", "parkjunwoo-quest"},
		{"parkjunwoo-quest 0.1", "parkjunwoo-quest"},
		{"bot\textra", "bot"},
		{"  spaced/1.0", ""},
		{"", ""},
	}
	for _, tt := range tests {
		if got := ProductToken(tt.in); got != tt.want {
			t.Errorf("ProductToken(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
