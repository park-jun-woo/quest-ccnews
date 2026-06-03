//ff:type feature=ingestion type=model
//ff:what 한 ingestion 트랙의 현재 위치와 실행 상태(running/waiting/exhausted).

package session

// Cursor: a single ingestion track position and run state.
type Cursor struct {
	Cursor string `json:"cursor"` // current WARC name
	State  string `json:"state"`  // running|waiting (forward) / running|exhausted (backward)
}
