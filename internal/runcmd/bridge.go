//ff:func feature=ingestion type=helper control=iteration dimension=1 level=error
//ff:what scratch session.Session → reins quest.Session 브리지. 새로 스캔된 기사를 reins Item(Key=URL, Payload=*session.Article, State=TODO)으로 append하되, 이미 시드된 URL은 건너뛴다(중복 방지). robots 거부 기사는 blockArticle로 BLOCKED 직접 시드(G4, 게이트 미경유, SkipReason 보존). 커서/processed→Meta["ingestion"], 호스트 robots 캐시→Meta["hosts"], UA→Meta["user_agent"](G2). 처리 뒤 scratch.Articles는 비워 같은 기사가 두 번 브리지되지 않게 한다. 새로 시드한 기사 수를 돌려준다.

package runcmd

import (
	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

// bridge folds the ingestion scratch into the reins session: it seeds each newly
// scanned article as a reins Item and persists the cursor/host cache into Meta.
//
// Dedup: an article whose URL already has an Item is skipped (the processed_warcs
// ratchet normally prevents re-scanning, but dedup keeps a re-run idempotent).
//
// Robots (G4): a denied article is seeded straight to BLOCKED via blockArticle —
// bypassing the submit gate — and its SkipReason is recorded on the payload.
// Allowed articles are seeded TODO for the agent to pick up.
//
// After folding, scratch.Articles is truncated so the next bridge call only sees
// the next WARC's articles. now is the timestamp Apply stamps onto BLOCKED items.
func bridge(scratch *session.Session, s *quest.Session, guard *robotsGuard, now string) int {
	seen := make(map[string]bool, len(s.Items))
	for _, it := range s.Items {
		seen[it.Key] = true
	}

	seeded := 0
	for _, a := range scratch.Articles {
		if a == nil || a.URL == "" || seen[a.URL] {
			continue
		}
		seen[a.URL] = true

		it := &quest.Item{Key: a.URL, State: quest.TODO}
		blockArticle(it, a, guard, now)
		// SetPayload is a snapshot, so it must run after blockArticle mutates
		// a.State/a.SkipReason — otherwise the deny reason is lost from the
		// persisted payload (and from export).
		if err := it.SetPayload(a); err != nil {
			return seeded
		}
		s.Items = append(s.Items, it)
		seeded++
	}
	scratch.Articles = scratch.Articles[:0]

	s.SetMeta(metaUserAgent, scratch.UserAgent)
	s.SetMeta(metaIngestion, scratch.Ingestion)
	s.SetMeta(metaHosts, scratch.Hosts)
	return seeded
}
