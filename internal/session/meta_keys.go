//ff:type feature=session type=model
//ff:what reins quest.Session.Meta(G2) 슬롯에 인제스천·재독 상태를 보존하는 키 상수(SSOT). runcmd(run: 쓰기)와 ccnewsquest(Prepare/Render: 읽기·robots 캐시 갱신)가 같은 키를 공유하도록 session 패키지에 둔다. user_agent(크롤 UA), ingestion(투트랙 커서/processed), hosts(호스트별 robots 캐시), cache_dir(WARC 캐시 절대경로 — Phase013 B).
package session

// Meta slot keys under reins quest.Session.Meta (G2). Shared between runcmd
// (the `run` ingestion writes them) and ccnewsquest (Prepare/Render read them, and
// Prepare updates the robots cache). Keeping them here makes the Meta schema a
// single source of truth across the two packages.
const (
	MetaUserAgent = "user_agent" // crawl User-Agent (string)
	MetaIngestion = "ingestion"  // two-track cursor + processed_warcs (session.Ingestion)
	MetaHosts     = "hosts"      // per-host robots cache (map[string]*session.Host)
	MetaCacheDir  = "cache_dir"  // absolute WARC cache directory (string) — Phase013 B
)
