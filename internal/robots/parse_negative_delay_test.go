//ff:func feature=robots type=helper control=sequence
//ff:what 음수 Crawl-delay 값을 Parse가 무시하고 HasDelay를 세우지 않는지 검증한다.

package robots

import "testing"

func TestParseNegativeCrawlDelayIgnored(t *testing.T) {
	content := []byte("User-agent: *\nCrawl-delay: -1\n")
	rs := Parse(content)
	if rs.Groups[0].HasDelay {
		t.Errorf("negative crawl-delay should not set HasDelay")
	}
}
