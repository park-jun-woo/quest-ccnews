//ff:func feature=ingestion type=helper control=sequence
//ff:what Run이 트랙이 하나도 없으면 아무 일도 안 하고(no-op) 끝나는지 검증한다.

package ingest

import (
	"bytes"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestRunNoTracks(t *testing.T) {
	c := NewClient("ua", t.TempDir())
	s := session.New("ua", "cc-news")
	if err := Run(c, s, RunOptions{Tracks: nil}, new(bytes.Buffer)); err != nil {
		t.Fatalf("Run with no tracks should be a no-op: %v", err)
	}
}
