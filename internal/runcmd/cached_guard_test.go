//ff:func feature=robots type=helper control=sequence
//ff:what cachedGuard 테스트 헬퍼. host의 ruleset 캐시를 주어진 rs로 미리 시드한 robotsGuard를 만들어 allowed()가 fetch 없이 캐시 적중 경로(steady state)만 결정적으로 타게 한다.

package runcmd

import (
	"github.com/park-jun-woo/quest-ccnews/internal/robots"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// cachedGuard returns a robotsGuard with host's ruleset pre-seeded to rs, so
// allowed() exercises the cache-hit (steady-state) path with no network fetch.
func cachedGuard(host string, rs *robots.Ruleset) *robotsGuard {
	g := newRobotsGuard("parkjunwoo-quest/0.1", map[string]*session.Host{})
	g.rulesets[host] = rs
	return g
}
