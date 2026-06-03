//ff:func feature=ingestion type=helper control=sequence level=error
//ff:what forward 트랙의 다음 WARC를 고른다. 최신(현재) 월만 본다 — 다 처리했으면 done=true(→waiting, 폴링). fetchPaths 주입으로 IO 분리.

package ingest

// SelectForward finds the next unprocessed WARC for the forward track. Forward
// only chases the newest dump, so it inspects a single month (the current one)
// and reports done=true when that month is fully processed — the engine then sets
// the track to StateWaiting to poll for new dumps. fetchPaths is injected for
// unit-testability (the engine passes Client.FetchPaths).
func SelectForward(current Month, processed map[string]bool, fetchPaths func(Month) ([]string, error)) (path string, done bool, err error) {
	paths, ferr := fetchPaths(current)
	if ferr != nil {
		return "", false, ferr
	}
	if p, ok := NextUnprocessed(paths, processed); ok {
		return p, false, nil
	}
	return "", true, nil
}
