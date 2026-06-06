//ff:func feature=ingestion type=helper control=iteration dimension=1 level=error
//ff:what scratch session.Session → reins quest.Session 브리지. 새로 스캔된 기사를 reins Item(Key=URL, Payload=*session.Article, State=TODO)으로 append하되, 이미 시드된 URL은 건너뛴다(중복 방지). robots는 시드 시 fetch하지 않고 전부 TODO로 적재한다(Phase013 A: robots 판정은 pick-time Prepare로 이동 — 대량 인제스천 시 호스트 라이브 fetch 폭주 회피). 커서/processed→Meta["ingestion"], 호스트 robots 캐시→Meta["hosts"], UA→Meta["user_agent"], WARC 캐시 절대경로→Meta["cache_dir"](Phase013 B). 처리 뒤 scratch.Articles는 비워 같은 기사가 두 번 브리지되지 않게 한다. 새로 시드한 기사 수를 돌려준다.
package runcmd

import (
	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

// bridge folds the ingestion scratch into the reins session: it seeds each newly
// scanned article as a reins Item and persists the cursor/host cache + cache dir
// into Meta.
//
// Dedup: an article whose URL already has an Item is skipped (the processed_warcs
// ratchet normally prevents re-scanning, but dedup keeps a re-run idempotent).
//
// Robots (Phase013 A): seeding no longer fetches robots.txt — every article is
// seeded TODO and the per-host robots decision is deferred to pick time (Prepare),
// where a denied host short-circuits to BLOCKED. This avoids the seed-time live
// fetch storm (≈1,892 hosts on a single CC-NEWS WARC).
//
// cacheDir is the absolute WARC cache directory recorded into Meta so submit/next
// re-read from the same path regardless of CWD (Phase013 B). After folding,
// scratch.Articles is truncated so the next bridge call only sees the next WARC's
// articles.
func bridge(scratch *session.Session, s *quest.Session, cacheDir string) int {
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
	if cacheDir != "" {
		s.SetMeta(metaCacheDir, cacheDir)
	}
	return seeded
}
