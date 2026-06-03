//ff:func feature=ingestion type=helper control=iteration dimension=1
//ff:what processed_warcs 슬라이스를 빠른 조회용 집합(map)으로 변환한다. 순수 함수.

package ingest

// ProcessedSet builds a lookup set from a processed_warcs slice so the ratchet
// lock can be checked in O(1). Pure (no IO).
func ProcessedSet(names []string) map[string]bool {
	set := make(map[string]bool, len(names))
	for _, n := range names {
		set[n] = true
	}
	return set
}
