//ff:func feature=ingestion type=helper control=sequence
//ff:what EnsureCursor가 backward 커서를 running 상태로 새로 만들어 붙이는지 검증한다.

package ingest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestEnsureCursorCreatesBackward(t *testing.T) {
	ing := &session.Ingestion{}
	c := EnsureCursor(ing, Backward)
	if c == nil || ing.Backward != c || c.State != StateRunning {
		t.Fatalf("backward cursor not created: %+v", c)
	}
}
