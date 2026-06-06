//ff:type feature=ingestion type=model
//ff:what reins quest.Session.Meta(G2) 키 상수의 runcmd 별칭. SSOT는 session.Meta* 상수(두 패키지 공유). user_agent·ingestion·hosts·cache_dir(WARC 캐시 절대경로, Phase013 B). run이 매 step 갱신해 다음 run·submit이 커서·캐시 경로를 이어받게 한다.
package runcmd

import "github.com/park-jun-woo/quest-ccnews/internal/session"

// Meta slot keys under quest.Session.Meta (G2), aliasing the session package's
// shared constants so runcmd and ccnewsquest agree on the schema. The ingestion
// run reads them on start (to resume cursors) and writes them back on every WARC
// so an interrupt is always resumable from the persisted reins session.
const (
	metaUserAgent = session.MetaUserAgent // crawl User-Agent (string)
	metaIngestion = session.MetaIngestion // two-track cursor + processed_warcs (session.Ingestion)
	metaHosts     = session.MetaHosts     // per-host robots cache (map[string]*session.Host)
	metaCacheDir  = session.MetaCacheDir  // absolute WARC cache directory (string) — Phase013 B
)
