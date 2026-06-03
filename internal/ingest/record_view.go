//ff:type feature=ingestion type=model
//ff:what 한 WARC response 레코드에서 순수 변환에 필요한 최소 정보(타입, Target-URI, 오프셋). IO 경계와 순수 로직을 가르는 값.

package ingest

// RecordView is the minimal, IO-free view of a single WARC record that the pure
// conversion logic needs. The thin WARC scanner fills it from a warc.Record; the
// pure converter (ToArticle) turns it into a session.Article without touching IO.
type RecordView struct {
	Type      string // WARC-Type header, e.g. "response"
	TargetURI string // WARC-Target-URI header
	Offset    int64  // logical locator within the WARC stream
}
