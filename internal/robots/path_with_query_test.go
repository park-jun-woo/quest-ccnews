//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what pathWithQuery가 URL에서 escaped path와 raw query를 합치고, query 없으면 path만, 파싱 실패 시 원본을 반환하는지 테이블로 검증한다.

package robots

import "testing"

func TestPathWithQuery(t *testing.T) {
	tests := []struct {
		name   string
		rawURL string
		want   string
	}{
		{"path only", "https://e.com/a/b", "/a/b"},
		{"path with query", "https://e.com/a?x=1&y=2", "/a?x=1&y=2"},
		{"root", "https://e.com/", "/"},
		{"escaped path preserved", "https://e.com/a%20b", "/a%20b"},
		{"parse error falls back", "://bad", "://bad"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathWithQuery(tt.rawURL); got != tt.want {
				t.Errorf("pathWithQuery(%q) = %q, want %q", tt.rawURL, got, tt.want)
			}
		})
	}
}
