//ff:func feature=ingestion type=helper control=selection
//ff:what --track 플래그 값(forward|backward|both)을 실행할 트랙 목록으로 매핑한다. 순수 함수.

package ingest

// TracksFor maps a --track flag value to the track list to run.
func TracksFor(flag string) []Track {
	switch flag {
	case "forward":
		return []Track{Forward}
	case "backward":
		return []Track{Backward}
	default: // "both"
		return []Track{Forward, Backward}
	}
}
