//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what MatchPattern이 빈 패턴·$ 앵커·* 와일드카드(선행/중간/다중)·prefix 매칭을 RFC 9309대로 판정하는지 테이블로 검증한다.

package robots

import "testing"

func TestMatchPattern(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		path    string
		want    bool
	}{
		{"empty pattern matches all", "", "/anything", true},
		{"empty pattern matches root", "", "/", true},
		{"dollar only matches empty path", "$", "", true},
		{"dollar only rejects non-empty", "$", "/", false},
		{"literal prefix match", "/private", "/private/x", true},
		{"literal prefix mismatch at head", "/private", "/public/private", false},
		{"exact literal", "/a", "/a", true},
		{"wildcard middle", "/a*c", "/abxc", true},
		{"wildcard middle no match", "/a*c", "/abx", false},
		{"trailing wildcard", "/a*", "/abcdef", true},
		{"leading wildcard then literal", "*foo", "/barfoo", true},
		{"leading wildcard literal not present", "*foo", "/barbaz", false},
		{"anchor end exact", "/a$", "/a", true},
		{"anchor end longer fails", "/a$", "/ab", false},
		{"anchor end with trailing wildcard", "/a*$", "/abc", true},
		{"wildcard plus anchor exact end", "/a*c$", "/abc", true},
		{"wildcard plus anchor not at end", "/a*c$", "/abcd", false},
		{"multiple wildcards in order", "/a*b*c", "/aXXbYYc", true},
		{"multiple wildcards out of order", "/a*b*c", "/aXXcYYb", false},
		{"prefix style allows suffix without anchor", "/dir", "/dir/sub/page", true},
		{"first segment must be at head", "/x", "y/x", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatchPattern(tt.pattern, tt.path); got != tt.want {
				t.Errorf("MatchPattern(%q, %q) = %v, want %v", tt.pattern, tt.path, got, tt.want)
			}
		})
	}
}
