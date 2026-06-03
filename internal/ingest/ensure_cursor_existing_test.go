//ff:func feature=ingestion type=helper control=sequence
//ff:what EnsureCursor가 이미 있는 forward/backward 커서를 그대로 반환하는지 검증한다.

package ingest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestEnsureCursorReturnsExisting(t *testing.T) {
	existingF := &session.Cursor{Cursor: "f.warc.gz", State: StateWaiting}
	existingB := &session.Cursor{Cursor: "b.warc.gz", State: StateExhausted}
	ing := &session.Ingestion{Forward: existingF, Backward: existingB}
	if EnsureCursor(ing, Forward) != existingF {
		t.Error("forward should return existing")
	}
	if EnsureCursor(ing, Backward) != existingB {
		t.Error("backward should return existing")
	}
}
