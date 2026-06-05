//ff:type feature=ingestion type=model
//ff:what reins quest.Session.Meta(G2)에 인제스천 상태를 보존하는 키 상수. user_agent(크롤 UA), ingestion(투트랙 커서/processed), hosts(호스트별 robots 캐시). run이 매 step 갱신해 다음 run에 커서가 이어지게 한다.

package runcmd

// Meta slot keys under quest.Session.Meta (G2). The ingestion run reads them on
// start (to resume cursors) and writes them back on every WARC so an interrupt is
// always resumable from the persisted reins session.
const (
	metaUserAgent = "user_agent" // crawl User-Agent (string)
	metaIngestion = "ingestion"  // two-track cursor + processed_warcs (session.Ingestion)
	metaHosts     = "hosts"      // per-host robots cache (map[string]*session.Host)
)
