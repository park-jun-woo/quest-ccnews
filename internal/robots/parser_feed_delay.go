//ff:func feature=robots type=helper control=sequence
//ff:what Crawl-delay 줄을 현재 그룹에 반영한다. 음이 아닌 정수만 받아 CrawlDelay/HasDelay를 설정한다.

package robots

import (
	"strconv"
	"strings"
)

// feedDelay applies a Crawl-delay line to the current group. A directive before
// any user-agent is ignored. Only a non-negative integer is accepted; it sets
// CrawlDelay and HasDelay (an explicit "0" still sets HasDelay). Ends the
// user-agent run.
func (p *parser) feedDelay(value string) {
	if p.cur == nil {
		return
	}
	p.expectingAgent = false
	n, err := strconv.Atoi(strings.TrimSpace(value))
	if err != nil || n < 0 {
		return
	}
	p.cur.CrawlDelay = n
	p.cur.HasDelay = true
}
