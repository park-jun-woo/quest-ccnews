//ff:type feature=robots type=model
//ff:func feature=robots type=helper control=sequence
//ff:what robotsGuard는 호스트 단위 robots 평가 캐시(G3). 호스트당 robots.txt를 1회 fetch·파싱해 여러 기사가 공유한다. robots 클라이언트, product token, 그리고 scratch.Hosts(robots 캐시 보존처)와 파싱된 ruleset 메모리 캐시를 묶는다. newRobotsGuard는 hosts를 보존 robots 캐시로 공유하는 guard를 만든다.

package runcmd

import (
	"github.com/park-jun-woo/quest-ccnews/internal/robots"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// robotsGuard evaluates articles against per-host robots rules, fetching and
// parsing each host's robots.txt at most once (G3 host-scoped shared cache). The
// session.Robots record is persisted into hosts (→ Session.Meta["hosts"]); the
// parsed *robots.Ruleset is held in-memory in rulesets for the run's lifetime.
type robotsGuard struct {
	client       *robots.Client
	productToken string
	hosts        map[string]*session.Host
	rulesets     map[string]*robots.Ruleset
}

// newRobotsGuard builds a guard sharing hosts as the persisted robots cache.
func newRobotsGuard(userAgent string, hosts map[string]*session.Host) *robotsGuard {
	return &robotsGuard{
		client:       robots.NewClient(userAgent),
		productToken: robots.ProductToken(userAgent),
		hosts:        hosts,
		rulesets:     map[string]*robots.Ruleset{},
	}
}
