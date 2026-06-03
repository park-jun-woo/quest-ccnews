//ff:func feature=ingestion type=helper control=sequence
//ff:what 한 WARC 처리 완료를 커서/processed_warcs에 반영한다(래칫 잠금). 커서를 처리한 WARC로 옮기고 상태를 running으로 둔다. 순수 함수.

package ingest

import "github.com/park-jun-woo/quest-ccnews/internal/session"

// AdvanceCursor records that warcName has been fully processed: it appends the
// name to processed_warcs (idempotent ratchet lock), points the cursor at it,
// and sets the track running. It mutates only the passed values (no IO).
func AdvanceCursor(ing *session.Ingestion, cur *session.Cursor, warcName string) {
	if !ProcessedSet(ing.ProcessedWarcs)[warcName] {
		ing.ProcessedWarcs = append(ing.ProcessedWarcs, warcName)
	}
	cur.Cursor = warcName
	cur.State = StateRunning
}
