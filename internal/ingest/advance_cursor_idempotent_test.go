//ff:func feature=ingestion type=helper control=sequence
//ff:what AdvanceCursor가 이미 processed에 있는 WARC를 중복 추가하지 않는지(래칫 멱등성) 검증한다.

package ingest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestAdvanceCursorIdempotent(t *testing.T) {
	ing := &session.Ingestion{ProcessedWarcs: []string{"a.warc.gz"}}
	cur := &session.Cursor{}

	AdvanceCursor(ing, cur, "a.warc.gz")

	if len(ing.ProcessedWarcs) != 1 {
		t.Fatalf("processed = %v, want no duplicate", ing.ProcessedWarcs)
	}
	if cur.Cursor != "a.warc.gz" || cur.State != StateRunning {
		t.Errorf("cursor=%q state=%q", cur.Cursor, cur.State)
	}
}
