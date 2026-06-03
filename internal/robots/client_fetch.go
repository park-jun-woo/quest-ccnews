//ff:func feature=robots type=helper control=selection level=error
//ff:what 호스트당 robots.txt 1회 fetch → 파싱 → session.Robots 캐시 레코드와 룰셋 반환. 2xx=ok, 4xx=missing(허용), 5xx/타임아웃=unreachable.

package robots

import (
	"net/http"
	"time"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// Fetch performs the one and only live network call of the pipeline: it GETs
// https://<host>/robots.txt once with our User-Agent header and maps the result
// to a session.Robots cache record plus the parsed Ruleset.
//
// Status mapping (Phase004 §robots규칙):
//   - 2xx                   → "ok"          (rules parsed and applied)
//   - 4xx (incl. 404)       → "missing"     (no robots = crawl allowed)
//   - 5xx / transport error → "unreachable" (conservative; gate decides policy)
//
// On "missing"/"unreachable" the returned Ruleset is empty (default-allow);
// callers reading CrawlAllowed and Status drive the conservative hold for
// "unreachable" per Phase004 §열린결정.
func (c *Client) Fetch(host string) (*session.Robots, *Ruleset, error) {
	rec := &session.Robots{
		FetchedAt: time.Now().UTC().Format(time.RFC3339),
		RobotsURL: RobotsURL(host),
	}

	req, err := http.NewRequest(http.MethodGet, rec.RobotsURL, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("User-Agent", c.userAgent)

	resp, err := c.http.Do(req)
	if err != nil {
		// Transport error / timeout → unreachable (conservative).
		rec.Status = "unreachable"
		rec.CrawlAllowed = false
		return rec, &Ruleset{}, nil
	}
	defer resp.Body.Close()

	switch {
	case resp.StatusCode >= 200 && resp.StatusCode < 300:
		return c.fetchOK(rec, resp)
	case resp.StatusCode >= 400 && resp.StatusCode < 500:
		rec.Status = "missing"
		rec.CrawlAllowed = true
		return rec, &Ruleset{}, nil
	default: // 5xx and anything else
		rec.Status = "unreachable"
		rec.CrawlAllowed = false
		return rec, &Ruleset{}, nil
	}
}
