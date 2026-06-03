//ff:type feature=ingestion type=model
//ff:what CC-NEWS WARC 파일을 훑는 양방향 ingestion 커서. forward/backward 트랙과 처리 완료 WARC 목록.

package session

// Ingestion: the two-track ingestion cursor over CC-NEWS WARC files.
type Ingestion struct {
	Source         string   `json:"source"`             // e.g. "cc-news"
	Forward        *Cursor  `json:"forward,omitempty"`  // newest → waiting
	Backward       *Cursor  `json:"backward,omitempty"` // past → exhausted
	ProcessedWarcs []string `json:"processed_warcs"`    // already-processed WARC names
}
