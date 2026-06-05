//ff:func feature=ingestion type=helper control=sequence level=error
//ff:what reins quest.Session.Meta(G2)에서 인제스천 스크래치 session.Session을 복원한다. user_agent·ingestion(커서/processed)·hosts(robots 캐시)를 디코드해 ingest 루프가 커서에서 재개하게 한다. Meta가 비면 defaultUA로 빈 스크래치를 만든다. Articles는 비워 둔다 — 이번 run에서 새로 스캔한 기사만 담는 통로다.

package runcmd

import (
	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

// restoreScratch rebuilds the in-memory ingestion scratch (a session.Session that
// the ingest loop mutates) from the reins session's Meta slot, so the two-track
// cursor and the per-host robots cache resume across runs. A missing Meta yields a
// fresh scratch seeded with defaultUA. Articles is intentionally left empty: the
// scratch's Articles slice is just the conduit for newly scanned articles this run;
// already-seeded items live in the reins session, not here.
func restoreScratch(s *quest.Session, defaultUA string) (*session.Session, error) {
	ua := defaultUA
	if v, ok := s.GetMeta(metaUserAgent); ok {
		if str, ok := v.(string); ok && str != "" {
			ua = str
		}
	}
	scratch := session.New(ua, "cc-news")

	if _, err := decodeMeta(s, metaIngestion, &scratch.Ingestion); err != nil {
		return nil, err
	}
	if _, err := decodeMeta(s, metaHosts, &scratch.Hosts); err != nil {
		return nil, err
	}
	if scratch.Hosts == nil {
		scratch.Hosts = map[string]*session.Host{}
	}
	if scratch.Ingestion.ProcessedWarcs == nil {
		scratch.Ingestion.ProcessedWarcs = []string{}
	}
	return scratch, nil
}
