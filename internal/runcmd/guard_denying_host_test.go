//ff:func feature=robots type=helper control=sequence
//ff:what guardDenyingHost 테스트 헬퍼. host의 ruleset 캐시를 "/" 전면 거부로 미리 시드한 robotsGuard를 만들어 allowed()가 네트워크 없이 캐시만으로 모든 URL을 거부하게 한다.

package runcmd

import (
	"github.com/park-jun-woo/quest-ccnews/internal/robots"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// guardDenyingHost returns a robotsGuard whose ruleset cache is pre-seeded for
// host so allowed() hits the cache (no network). The pre-seeded ruleset disallows
// everything under "/", so any URL on that host is denied.
func guardDenyingHost(host string) *robotsGuard {
	g := newRobotsGuard("parkjunwoo-quest/0.1", map[string]*session.Host{})
	g.rulesets[host] = &robots.Ruleset{Groups: []robots.Group{{
		Agents: []string{"*"},
		Rules:  []robots.Rule{{Allow: false, Pattern: "/"}},
	}}}
	return g
}
