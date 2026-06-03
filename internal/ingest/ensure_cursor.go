//ff:func feature=ingestion type=helper control=selection
//ff:what 트랙(forward/backward)의 커서를 반환하되, 아직 없으면 running 상태로 새로 만들어 붙인다. 순수 함수.

package ingest

import "github.com/park-jun-woo/quest-ccnews/internal/session"

// EnsureCursor returns the cursor for the given track, creating a fresh running
// cursor on the ingestion if the track has none yet (e.g. a brand-new session).
// Mutates ing only (no IO).
func EnsureCursor(ing *session.Ingestion, track Track) *session.Cursor {
	switch track {
	case Backward:
		if ing.Backward == nil {
			ing.Backward = &session.Cursor{State: StateRunning}
		}
		return ing.Backward
	default:
		if ing.Forward == nil {
			ing.Forward = &session.Cursor{State: StateRunning}
		}
		return ing.Forward
	}
}
