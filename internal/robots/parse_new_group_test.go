//ff:func feature=robots type=helper control=sequence
//ff:what 규칙 뒤에 오는 user-agent 줄이 새 그룹을 시작하는지 Parse로 검증한다.

package robots

import "testing"

func TestParseNewGroupAfterRule(t *testing.T) {
	content := []byte("User-agent: a\nDisallow: /x\nUser-agent: b\nDisallow: /y\n")
	rs := Parse(content)
	if len(rs.Groups) != 2 {
		t.Fatalf("groups = %d, want 2", len(rs.Groups))
	}
}
