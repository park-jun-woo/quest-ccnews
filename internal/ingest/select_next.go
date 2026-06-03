//ff:func feature=ingestion type=helper control=iteration dimension=1 level=error
//ff:what 한 트랙에서 처리할 다음 WARC 경로를 월을 거슬러 가며 찾는다. fetchPaths를 주입받아 IO와 분리(테스트 가능). 더 받을 게 없으면 done=true.

package ingest

// SelectNext finds the next unprocessed WARC object path for a track, walking
// months from `start` toward the past. For each month it asks fetchPaths for the
// listing and picks the newest unprocessed WARC (NextUnprocessed). When a month is
// fully processed it steps to the previous month (Month.Prev); reaching the
// earliest CC-NEWS dump with nothing left yields done=true (no more WARCs).
//
// fetchPaths is injected so this stepper is unit-testable without network IO; the
// engine passes Client.FetchPaths. processed is the ratchet lock set.
func SelectNext(start Month, processed map[string]bool, fetchPaths func(Month) ([]string, error)) (path string, done bool, err error) {
	m := start
	for {
		paths, ferr := fetchPaths(m)
		if ferr != nil {
			return "", false, ferr
		}
		if p, ok := NextUnprocessed(paths, processed); ok {
			return p, false, nil
		}
		prev, ok := m.Prev()
		if !ok {
			return "", true, nil
		}
		m = prev
	}
}
