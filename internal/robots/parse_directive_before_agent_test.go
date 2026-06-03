//ff:func feature=robots type=helper control=sequence
//ff:what user-agent 이전에 나온 Disallow/Crawl-delay 지시자를 Parse가 무시하는지 검증한다.

package robots

import "testing"

func TestParseDirectiveBeforeAnyAgentIgnored(t *testing.T) {
	content := []byte("Disallow: /x\nCrawl-delay: 3\nUser-agent: *\nAllow: /\n")
	rs := Parse(content)
	if len(rs.Groups) != 1 {
		t.Fatalf("groups = %d, want 1", len(rs.Groups))
	}
	if len(rs.Groups[0].Rules) != 1 || !rs.Groups[0].Rules[0].Allow {
		t.Errorf("expected only the Allow rule, got %+v", rs.Groups[0].Rules)
	}
	if rs.Groups[0].HasDelay {
		t.Errorf("crawl-delay before agent should be ignored")
	}
}
