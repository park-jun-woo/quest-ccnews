//ff:func feature=robots type=helper control=sequence
//ff:what robots 게이트 연동. 기사 path를 룰셋에 평가해 거부면 Article을 BLOCKED로 잠그고 skip_reason 기록. 허용이면 통과(true 반환).

package robots

import (
	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// Gate applies the deterministic robots decision to one article. It evaluates
// the article's URL against the host ruleset for productToken; on deny it locks
// the article to BLOCKED with skip_reason "robots Disallow: <rule>" and returns
// false. On allow it leaves the article untouched and returns true so the
// caller proceeds to the next gate (Phase005). State mutation only — no IO.
//
// Only TODO articles are gated; an already-locked article is left as-is.
func Gate(a *session.Article, rs *Ruleset, productToken string) (allowed bool) {
	if a.State != session.TODO {
		return a.State != session.BLOCKED
	}
	d := Evaluate(rs, productToken, a.URL)
	if d.Allowed {
		return true
	}
	a.State = session.BLOCKED
	a.SkipReason = "robots " + d.Rule
	return false
}
