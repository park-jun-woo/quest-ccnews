//ff:type feature=ingestion type=model
//ff:what ingestion 트랙 식별 타입(Track)과 트랙 상태 문자열 상수(running/waiting/exhausted).

package ingest

// Track run states stored in session.Cursor.State.
const (
	StateRunning   = "running"   // actively processing WARCs
	StateWaiting   = "waiting"   // forward only: caught up, polling for new dumps
	StateExhausted = "exhausted" // backward only: reached the earliest CC-NEWS dump
)

// Track identifies an ingestion direction.
type Track string

const (
	Forward  Track = "forward"
	Backward Track = "backward"
)
