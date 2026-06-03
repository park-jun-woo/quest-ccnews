//ff:func feature=robots type=helper control=sequence
//ff:what 2xx 응답 본문을 1 MiB까지 읽어 룰셋으로 파싱하고, 우리 그룹의 crawl-delay를 ok 캐시 레코드에 채운다.

package robots

import (
	"io"
	"net/http"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// fetchOK handles a 2xx robots.txt response: it reads the body (capped at 1 MiB
// — robots.txt is small), parses it into a Ruleset, marks the record "ok" and
// crawl-allowed, and copies the crawl-delay from the group that applies to our
// product token when one is set.
func (c *Client) fetchOK(rec *session.Robots, resp *http.Response) (*session.Robots, *Ruleset, error) {
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 1<<20)) // robots.txt cap: 1 MiB
	rs := Parse(body)
	rec.Status = "ok"
	rec.CrawlAllowed = true
	if g := SelectGroup(rs, c.productToken); g != nil && g.HasDelay {
		rec.CrawlDelaySec = g.CrawlDelay
	}
	return rec, rs, nil
}
