//ff:func feature=robots type=helper control=sequence
//ff:what user-agent 사이의 알 수 없는 필드가 agent run을 끊어 다음 user-agent가 새 그룹을 시작하는지 Parse로 검증한다.

package robots

import "testing"

func TestParseUnknownFieldEndsAgentRun(t *testing.T) {
	// An unknown field between two user-agent lines must break the agent run,
	// so the second user-agent starts a new group.
	content := []byte("User-agent: a\nSitemap: https://x\nUser-agent: b\nDisallow: /y\n")
	rs := Parse(content)
	if len(rs.Groups) != 2 {
		t.Fatalf("groups = %d, want 2 (unknown field breaks run)", len(rs.Groups))
	}
}
