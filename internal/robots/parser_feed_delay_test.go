//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what feedDelay가 음수/비정수는 무시하고, 음이 아닌 정수(0 포함)만 CrawlDelay/HasDelay로 반영하며 agent run을 끝내는지 테이블로 검증한다.

package robots

import "testing"

func TestParserFeedDelay(t *testing.T) {
	tests := []struct {
		name         string
		value        string
		wantDelay    int
		wantHasDelay bool
	}{
		{"positive int", "5", 5, true},
		{"explicit zero", "0", 0, true},
		{"trims whitespace", "  3 ", 3, true},
		{"negative ignored", "-1", 0, false},
		{"non-integer ignored", "abc", 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{rs: &Ruleset{Groups: []Group{{}}}, expectingAgent: true}
			p.cur = &p.rs.Groups[0]
			p.feedDelay(tt.value)
			if p.cur.CrawlDelay != tt.wantDelay || p.cur.HasDelay != tt.wantHasDelay || p.expectingAgent {
				t.Errorf("feedDelay(%q) -> delay=%d hasDelay=%v expecting=%v, want delay=%d hasDelay=%v expecting=false",
					tt.value, p.cur.CrawlDelay, p.cur.HasDelay, p.expectingAgent, tt.wantDelay, tt.wantHasDelay)
			}
		})
	}

	t.Run("no group ignored", func(t *testing.T) {
		p := &parser{rs: &Ruleset{}}
		p.feedDelay("5") // must not panic and leaves nil cur
		if p.cur != nil {
			t.Errorf("cur should remain nil")
		}
	})
}
