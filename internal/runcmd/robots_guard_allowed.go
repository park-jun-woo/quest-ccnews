//ff:func feature=robots type=helper control=selection level=error
//ff:what 기사 한 건의 robots 허용 여부를 평가한다. 호스트의 ruleset을 캐시(없으면 1회 fetch·파싱)해 robots.Evaluate로 판정하고, 거부면 (false, "robots <rule>") 사유를 돌려준다. fetch 실패는 보수적으로 거부(unreachable). 호스트 robots 캐시는 scratch.Hosts에 보존돼 Meta로 직렬화된다.

package runcmd

import (
	"github.com/park-jun-woo/quest-ccnews/internal/robots"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// allowed reports whether the article URL on host is crawlable under that host's
// robots.txt. It fetches and parses the host's robots.txt at most once per run
// (cached in g.rulesets and g.hosts), then evaluates the URL. On deny it returns
// allowed=false plus a "robots <rule>" reason for the SkipReason. A fetch transport
// failure is conservative: the host is recorded unreachable and the article denied.
func (g *robotsGuard) allowed(host, url string) (ok bool, reason string) {
	rs, cached := g.rulesets[host]
	if !cached {
		rec, parsed, err := g.client.Fetch(host)
		if err != nil {
			// Network/IO error fetching robots.txt → conservative deny.
			g.rulesets[host] = &robots.Ruleset{}
			return false, "robots unreachable"
		}
		rs = parsed
		g.rulesets[host] = rs
		if g.hosts[host] == nil {
			g.hosts[host] = &session.Host{Robots: rec}
		} else {
			g.hosts[host].Robots = rec
		}
		if rec.Status == "unreachable" {
			return false, "robots unreachable"
		}
	}
	d := robots.Evaluate(rs, g.productToken, url)
	if d.Allowed {
		return true, ""
	}
	return false, "robots " + d.Rule
}
