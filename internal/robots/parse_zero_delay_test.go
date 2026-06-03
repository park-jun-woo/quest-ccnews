//ff:func feature=robots type=helper control=sequence
//ff:what 명시적 Crawl-delay: 0이 HasDelay를 세우고 값 0을 저장하는지 Parse로 검증한다.

package robots

import "testing"

func TestParseZeroCrawlDelaySetsHasDelay(t *testing.T) {
	content := []byte("User-agent: *\nCrawl-delay: 0\n")
	rs := Parse(content)
	if !rs.Groups[0].HasDelay || rs.Groups[0].CrawlDelay != 0 {
		t.Errorf("explicit 0 should set HasDelay with value 0: %+v", rs.Groups[0])
	}
}
