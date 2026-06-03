//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what NormalizePath가 bare path/빈 값/full URL/query/malformed를 평가용 path로 정규화하는지 테이블로 검증한다.

package robots

import "testing"

func TestNormalizePath(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"bare path passthrough", "/foo/bar", "/foo/bar"},
		{"empty becomes root", "", "/"},
		{"bare path without leading slash", "foo", "/foo"},
		{"full url strips scheme and host", "https://example.com/a/b", "/a/b"},
		{"full url with query", "https://example.com/a?x=1&y=2", "/a?x=1&y=2"},
		{"full url root path", "https://example.com", "/"},
		{"full url with escaped path", "https://example.com/a%20b", "/a%20b"},
		{"http scheme", "http://example.com/p", "/p"},
		{"malformed url falls back to raw with leading slash", "https://%zz", "/https://%zz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizePath(tt.in); got != tt.want {
				t.Errorf("NormalizePath(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
