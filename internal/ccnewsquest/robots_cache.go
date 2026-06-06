//ff:type feature=robots type=model
//ff:func feature=robots type=helper control=sequence
//ff:what pick-time robots 평가 캐시(Phase013 A). 시드(bridge)는 더 이상 robots를 fetch하지 않고, 기사를 실제로 제출할 때(Prepare) 그 host의 robots를 host당 1회 fetch·파싱해 in-memory rulesets에 캐시한다(2번째 기사부터 캐시 히트, fetch 0). ccnewsDef가 Def()에서 만든 이 캐시 포인터를 공유해 한 프로세스 동안 호스트당 1회만 fetch한다. fetch 함수는 robotsFetch 패키지 변수로 주입 가능(테스트 카운터). 평가 결과 session.Robots 레코드는 Session.Meta의 호스트 캐시(metaHosts)에 적재돼 저장/복원 라운드트립으로 보존된다.
package ccnewsquest

import (
	"sync"

	"github.com/park-jun-woo/quest-ccnews/internal/robots"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// robotsFetch is the one live network call, indirected through a package var so a
// test can substitute a no-network stub (and count invocations) to assert that the
// pick-time evaluation fetches each host's robots.txt at most once. Defaults to the
// real per-UA robots client fetch.
var robotsFetch = func(userAgent, host string) (*session.Robots, *robots.Ruleset, error) {
	return robots.NewClient(userAgent).Fetch(host)
}

// robotsCache holds the per-host parsed rulesets for one process so the pick-time
// robots evaluation fetches each host at most once. It is created in Def and shared
// by value-receiver ccnewsDef through a pointer, so every Prepare in the process
// sees the same cache. The persisted half of the cache (the session.Robots record)
// lives in quest.Session.Meta[metaHosts]; this in-memory half holds the parsed
// rules needed to evaluate further URLs without a refetch.
type robotsCache struct {
	mu       sync.Mutex
	rulesets map[string]*robots.Ruleset
}

// newRobotsCache builds an empty pick-time robots cache.
func newRobotsCache() *robotsCache {
	return &robotsCache{rulesets: map[string]*robots.Ruleset{}}
}
