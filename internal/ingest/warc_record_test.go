//ff:type feature=ingestion type=helper
//ff:what 테스트 헬퍼 타입. writeWarc에 넘길 한 레코드의 WARC-Type과 WARC-Target-URI.

package ingest

// warcRecord is one record's WARC-Type and WARC-Target-URI for writeWarc.
type warcRecord struct {
	Type string
	URI  string
}
