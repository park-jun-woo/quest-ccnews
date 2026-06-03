//ff:type feature=host type=model
//ff:what 현재 robots.txt 1회 재확인 결과. 크롤 허용 여부와 crawl-delay를 캐싱한다.

package session

// Robots: result of the one-time current robots.txt re-check.
type Robots struct {
	FetchedAt     string `json:"fetched_at"`
	RobotsURL     string `json:"robots_url"`
	Status        string `json:"status"`          // ok|missing(=allowed)|unreachable
	CrawlAllowed  bool   `json:"crawl_allowed"`   // crawl-allowed (user-specified)
	CrawlDelaySec int    `json:"crawl_delay_sec"` // crawl-delay seconds
}
