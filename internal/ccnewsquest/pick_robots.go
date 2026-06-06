//ff:func feature=robots type=helper control=sequence level=error
//ff:what 한 기사의 robots 허용 여부를 pick-time(Prepare)에 평가한다(Phase013 A). host가 비면 허용(robots 무관). in-memory ruleset 캐시 적중이면 fetch 없이 평가; 미스면 robotsFetch로 host당 1회 fetch·파싱해 캐시하고 session.Robots 레코드를 Session.Meta의 호스트 캐시(metaHosts)에 적재해 라운드트립 보존한다. fetch 실패는 보수적 거부("robots unreachable"). robots.Evaluate로 URL 판정 → 거부면 (false, "robots <rule>"). Meta 쓰기는 submit이 Save하는 경로에서만 일어나므로 적재도 여기(Prepare 경로)에 둔다.
package ccnewsquest

import (
	"encoding/json"

	"github.com/park-jun-woo/quest-ccnews/internal/robots"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

// allowed evaluates whether article a may be crawled under its host's robots.txt,
// fetching that host at most once per process (in-memory ruleset cache) and recording
// the fetch result into the session Meta host cache so it survives the Save/Load
// round-trip. An empty host is allowed (no robots applies). A fetch transport failure
// is conservative: the host is denied "robots unreachable". On deny it returns
// allowed=false plus a "robots <rule>" reason for the SkipReason; on allow, ("", true).
func (c *robotsCache) allowed(s *quest.Session, userAgent string, a *session.Article) (ok bool, reason string) {
	if a.Host == "" {
		return true, ""
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	rs, cached := c.rulesets[a.Host]
	if !cached {
		rec, parsed, err := robotsFetch(userAgent, a.Host)
		if err != nil {
			// Transport/URL error → conservative deny; cache an empty ruleset so the
			// failed host is not re-fetched this process.
			c.rulesets[a.Host] = &robots.Ruleset{}
			return false, "robots unreachable"
		}
		rs = parsed
		c.rulesets[a.Host] = rs
		putHostRobots(s, a.Host, rec)
		if rec != nil && rec.Status == "unreachable" {
			return false, "robots unreachable"
		}
	}

	d := robots.Evaluate(rs, robots.ProductToken(userAgent), a.URL)
	if d.Allowed {
		return true, ""
	}
	return false, "robots " + d.Rule
}

// putHostRobots records the fetched robots record under host in the session Meta host
// cache (metaHosts), preserving any other hosts already cached. After quest.Load a
// Meta value is a generic map[string]any, so it is JSON round-tripped back into the
// typed map[string]*session.Host before the update, then written back via SetMeta so
// submit's Save persists it. Best-effort: a decode error simply starts a fresh map.
func putHostRobots(s *quest.Session, host string, rec *session.Robots) {
	hosts := map[string]*session.Host{}
	if v, ok := s.GetMeta(session.MetaHosts); ok {
		if b, err := json.Marshal(v); err == nil {
			_ = json.Unmarshal(b, &hosts)
		}
	}
	if hosts[host] == nil {
		hosts[host] = &session.Host{Robots: rec}
	} else {
		hosts[host].Robots = rec
	}
	s.SetMeta(session.MetaHosts, hosts)
}
