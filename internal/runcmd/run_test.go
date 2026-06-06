//ff:func feature=ingestion type=command control=sequence
//ff:what run 단위테스트(세션 파일 부재). 세션 파일이 없으면 빈 세션으로 시작·실행하고 최종 flush가 항상 저장하므로 파일이 생긴다(ingestRun 스텁으로 무네트워크).

package runcmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/ingest"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestRunMissingSessionStartsFresh(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "session.json")
	o, cmd := newOptions(t, path)
	stubIngest(t, func(*session.Session, ingest.RunOptions) error { return nil })

	if err := o.run(cmd, nil); err != nil {
		t.Fatalf("run() err = %v", err)
	}
	// Final flush always saves the session, so the file now exists.
	if _, err := os.Stat(path); err != nil {
		t.Errorf("session not saved: %v", err)
	}
}
