//ff:func feature=robots type=helper control=sequence
//ff:what allowed 단위테스트. 캐시 적중 시 빈 ruleset은 기본허용(ok), disallow 규칙은 "robots <rule>" 사유로 거부, 같은 캐시에서 형제 허용 경로는 네트워크 없이 허용; 캐시 미스+fetch 실패(URL 파싱 불가 호스트)는 보수적 거부 "robots unreachable"이며 빈 ruleset을 캐시해 재fetch하지 않음을 검증한다(fetch 성공 경로는 실네트워크라 tsma 최소 제외).

package runcmd

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/robots"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestAllowed(t *testing.T) {
	t.Run("cache hit, rule allows → ok", func(t *testing.T) {
		// Empty ruleset → no matching group → default-allow.
		g := cachedGuard("ex.com", &robots.Ruleset{})
		ok, reason := g.allowed("ex.com", "https://ex.com/article")
		if !ok || reason != "" {
			t.Errorf("ok=%v reason=%q, want true,\"\"", ok, reason)
		}
	})

	t.Run("cache hit, rule disallows → denied with rule reason", func(t *testing.T) {
		rs := &robots.Ruleset{Groups: []robots.Group{{
			Agents: []string{"*"},
			Rules:  []robots.Rule{{Allow: false, Pattern: "/private/"}},
		}}}
		g := cachedGuard("ex.com", rs)
		ok, reason := g.allowed("ex.com", "https://ex.com/private/x")
		if ok {
			t.Errorf("ok = true, want false (disallowed)")
		}
		if reason == "" || reason[:7] != "robots " {
			t.Errorf("reason = %q, want \"robots <rule>\"", reason)
		}
	})

	t.Run("fetch failure → conservative deny, ruleset cached empty", func(t *testing.T) {
		// A host with a space makes http.NewRequest fail at URL-parse time, so
		// Fetch returns an error WITHOUT any network call — deterministically
		// exercising allowed()'s cache-miss + fetch-error branch.
		g := newRobotsGuard("parkjunwoo-quest/0.1", map[string]*session.Host{})
		ok, reason := g.allowed("bad host", "https://bad host/x")
		if ok {
			t.Errorf("ok = true, want false (fetch error → conservative deny)")
		}
		if reason != "robots unreachable" {
			t.Errorf("reason = %q, want \"robots unreachable\"", reason)
		}
		// Empty ruleset is cached so the failed host is not refetched.
		if g.rulesets["bad host"] == nil {
			t.Errorf("failed host ruleset not cached")
		}
		// Second call hits the cache (default-allow on the empty ruleset).
		if ok2, _ := g.allowed("bad host", "https://bad host/x"); !ok2 {
			t.Errorf("second call did not hit cached empty ruleset")
		}
	})

	t.Run("cache hit does not consult network for an allowed sibling path", func(t *testing.T) {
		rs := &robots.Ruleset{Groups: []robots.Group{{
			Agents: []string{"*"},
			Rules:  []robots.Rule{{Allow: false, Pattern: "/private/"}},
		}}}
		g := cachedGuard("ex.com", rs)
		// /public is not under the disallow → allowed, purely from cache.
		ok, reason := g.allowed("ex.com", "https://ex.com/public/y")
		if !ok || reason != "" {
			t.Errorf("ok=%v reason=%q, want true,\"\"", ok, reason)
		}
	})
}
