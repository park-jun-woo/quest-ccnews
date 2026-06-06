//ff:func feature=robots type=helper control=sequence level=error
//ff:what robotsCache.allowed 단위 테스트. pick-time robots 허용 판정의 분기를 직접 커버한다. ① host 빈 문자열 → fetch 없이 허용. ② 허용 ruleset → (true,""). ③ 거부 ruleset → (false,"robots <rule>"). ④ fetch 에러 → 보수적 거부("robots unreachable")·빈 ruleset 캐시. ⑤ 같은 host 두 번 → fetch 1회(in-memory 캐시 히트). ⑥ fetch 성공 시 session.Robots 레코드가 Meta에 적재.
package ccnewsquest

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/robots"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestRobotsCacheAllowed(t *testing.T) {
	t.Run("empty host → allowed without fetch", func(t *testing.T) {
		calls := stubFetch(t, &session.Robots{Status: "ok"}, &robots.Ruleset{})
		c := newRobotsCache()
		ok, reason := c.allowed(quest.New(), "ua", &session.Article{Host: ""})
		if !ok || reason != "" {
			t.Fatalf("got (%v,%q), want (true,\"\")", ok, reason)
		}
		if *calls != 0 {
			t.Errorf("fetch calls = %d, want 0 (empty host needs no robots)", *calls)
		}
	})

	t.Run("allow ruleset → (true, \"\")", func(t *testing.T) {
		stubFetch(t, &session.Robots{Status: "ok"}, &robots.Ruleset{})
		c := newRobotsCache()
		ok, reason := c.allowed(quest.New(), "ua", &session.Article{Host: "a.com", URL: "https://a.com/x"})
		if !ok || reason != "" {
			t.Fatalf("got (%v,%q), want (true,\"\")", ok, reason)
		}
	})

	t.Run("deny ruleset → (false, robots <rule>)", func(t *testing.T) {
		stubFetch(t, &session.Robots{Status: "ok"}, denyRuleset())
		c := newRobotsCache()
		ok, reason := c.allowed(quest.New(), "ua", &session.Article{Host: "d.com", URL: "https://d.com/x"})
		if ok {
			t.Fatalf("got allowed=true, want false (deny ruleset)")
		}
		if !strings.HasPrefix(reason, "robots ") {
			t.Fatalf("reason = %q, want a \"robots <rule>\" deny reason", reason)
		}
	})

	t.Run("fetch error → conservative deny + empty ruleset cached", func(t *testing.T) {
		calls := 0
		orig := robotsFetch
		robotsFetch = func(string, string) (*session.Robots, *robots.Ruleset, error) {
			calls++
			return nil, nil, errors.New("transport down")
		}
		t.Cleanup(func() { robotsFetch = orig })

		c := newRobotsCache()
		ok, reason := c.allowed(quest.New(), "ua", &session.Article{Host: "fail.com", URL: "https://fail.com/x"})
		if ok || reason != "robots unreachable" {
			t.Fatalf("got (%v,%q), want (false,\"robots unreachable\")", ok, reason)
		}
		if c.rulesets["fail.com"] == nil {
			t.Errorf("failed host not cached; it would be re-fetched this process")
		}
		// A second call must not re-fetch the failed host.
		c.allowed(quest.New(), "ua", &session.Article{Host: "fail.com", URL: "https://fail.com/y"})
		if calls != 1 {
			t.Errorf("fetch calls = %d, want 1 (failed host cached, not re-fetched)", calls)
		}
	})

	t.Run("fetch ok but rec Status unreachable → deny", func(t *testing.T) {
		// No transport error, but the robots record itself reports the host as
		// unreachable (e.g. a 5xx). That is a conservative deny too (line 38).
		stubFetch(t, &session.Robots{Status: "unreachable"}, &robots.Ruleset{})
		c := newRobotsCache()
		ok, reason := c.allowed(quest.New(), "ua", &session.Article{Host: "u.com", URL: "https://u.com/x"})
		if ok || reason != "robots unreachable" {
			t.Fatalf("got (%v,%q), want (false,\"robots unreachable\")", ok, reason)
		}
	})

	t.Run("same host twice → fetched exactly once (cache hit)", func(t *testing.T) {
		calls := stubFetch(t, &session.Robots{Status: "ok"}, &robots.Ruleset{})
		c := newRobotsCache()
		s := quest.New()
		for i := 0; i < 2; i++ {
			c.allowed(s, "ua", &session.Article{Host: "h.com", URL: "https://h.com/x"})
		}
		if *calls != 1 {
			t.Errorf("fetch calls = %d, want 1 (second is a cache hit)", *calls)
		}
	})

	t.Run("fetch result recorded into Meta host cache", func(t *testing.T) {
		stubFetch(t, &session.Robots{Status: "ok", RobotsURL: "https://m.com/robots.txt"}, &robots.Ruleset{})
		c := newRobotsCache()
		s := quest.New()
		c.allowed(s, "ua", &session.Article{Host: "m.com", URL: "https://m.com/x"})

		v, ok := s.GetMeta(session.MetaHosts)
		if !ok {
			t.Fatal("Meta[hosts] absent after allow")
		}
		hosts := map[string]*session.Host{}
		b, _ := json.Marshal(v)
		_ = json.Unmarshal(b, &hosts)
		if hosts["m.com"] == nil || hosts["m.com"].Robots == nil {
			t.Fatalf("m.com robots not recorded into Meta: %#v", hosts)
		}
	})
}
