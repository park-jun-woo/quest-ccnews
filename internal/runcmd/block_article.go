//ff:func feature=ingestion type=helper control=sequence level=error
//ff:what robots 거부 기사를 BLOCKED로 직접 시드한다(G4, 게이트 미경유). guard가 있고 호스트가 있으면 guard.allowed로 평가하고, 거부면 payload.State/SkipReason을 보존한 뒤 quest.Apply(OutBlock)로 Item을 BLOCKED 시드한다. guard 부재·호스트 부재·허용이면 아무것도 하지 않아 호출자가 TODO로 둔다.

package runcmd

import (
	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

// blockArticle consults the robots guard for article a and, on a deny, seeds the
// reins Item straight to BLOCKED via quest.Apply with an OutBlock verdict —
// bypassing the submit gate (G4) — while preserving the deny reason on the payload.
// When there is no guard, the article has no host, or the host allows the URL, it
// leaves the Item untouched so the caller keeps it TODO. now is the Apply stamp.
func blockArticle(it *quest.Item, a *session.Article, guard *robotsGuard, now string) {
	if guard == nil || a.Host == "" {
		return
	}
	ok, reason := guard.allowed(a.Host, a.URL)
	if ok {
		return
	}
	a.State = session.BLOCKED
	a.SkipReason = reason
	quest.Apply(it, quest.Verdict{
		Outcome: quest.OutBlock,
		Facts:   []quest.Fact{{Rule: "robots", Where: a.Host, Actual: reason}},
	}, now)
}
