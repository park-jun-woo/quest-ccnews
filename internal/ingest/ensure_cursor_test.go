//ff:func feature=ingestion type=helper control=sequence
//ff:what EnsureCursor가 forward 커서를 새로 만들어 붙이고, 미지의 트랙은 forward로 매핑하는지 검증한다.

package ingest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestEnsureCursorCreatesForward(t *testing.T) {
	ing := &session.Ingestion{}
	c := EnsureCursor(ing, Forward)
	if c == nil || ing.Forward != c {
		t.Fatal("forward cursor not created/attached")
	}
	if c.State != StateRunning {
		t.Errorf("state = %q, want running", c.State)
	}
	// default branch (unknown track) maps to forward
	ing2 := &session.Ingestion{}
	if got := EnsureCursor(ing2, Track("weird")); got != ing2.Forward {
		t.Error("unknown track should use forward")
	}
}
