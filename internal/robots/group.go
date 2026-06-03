//ff:type feature=robots type=model
//ff:what RFC 9309 user-agent 그룹. 공유 agent 토큰들과 Allow/Disallow 규칙, crawl-delay를 담는다(순수 데이터).

package robots

// Group is one or more user-agent lines sharing a set of rules (an RFC 9309
// group). Agents holds lowercased product tokens; "*" is the catch-all.
type Group struct {
	Agents     []string
	Rules      []Rule
	CrawlDelay int  // seconds; 0 = unspecified
	HasDelay   bool // distinguishes an explicit "crawl-delay: 0" from absence
}
