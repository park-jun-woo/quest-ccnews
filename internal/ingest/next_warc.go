//ff:func feature=ingestion type=helper control=iteration dimension=1
//ff:what 한 달의 WARC 경로 목록에서 아직 처리하지 않은(processed_warcs에 없는) 다음 WARC를 newest→ 순으로 고른다. 순수 함수.

package ingest

// NextUnprocessed picks the next WARC object path to process from one month's
// warc.paths listing, scanning newest-first (CC-NEWS lists WARCs oldest-first by
// filename timestamp, so we walk the slice in reverse). A path is skipped if its
// basename is already in the processed set (ratchet lock). Returns ok=false when
// every WARC in the listing is already processed. Pure (no IO).
func NextUnprocessed(paths []string, processed map[string]bool) (path string, ok bool) {
	for i := len(paths) - 1; i >= 0; i-- {
		if !processed[WarcName(paths[i])] {
			return paths[i], true
		}
	}
	return "", false
}
