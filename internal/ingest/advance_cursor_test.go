//ff:func feature=ingestion type=helper control=sequence
//ff:what AdvanceCursor가 새 WARC를 processed에 추가하고 커서를 옮기며 상태를 running으로 만드는지 검증한다.

package ingest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestAdvanceCursorAppendsAndRatchets(t *testing.T) {
	ing := &session.Ingestion{ProcessedWarcs: []string{"a.warc.gz"}}
	cur := &session.Cursor{State: StateWaiting}

	AdvanceCursor(ing, cur, "b.warc.gz")

	if len(ing.ProcessedWarcs) != 2 || ing.ProcessedWarcs[1] != "b.warc.gz" {
		t.Fatalf("processed = %v, want [a.warc.gz b.warc.gz]", ing.ProcessedWarcs)
	}
	if cur.Cursor != "b.warc.gz" {
		t.Errorf("cursor = %q, want b.warc.gz", cur.Cursor)
	}
	if cur.State != StateRunning {
		t.Errorf("state = %q, want %q", cur.State, StateRunning)
	}
}
