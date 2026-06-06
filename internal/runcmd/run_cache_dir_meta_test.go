//ff:func feature=ingestion type=command control=sequence
//ff:what run 단위테스트(Phase013 B). run이 상대 cacheDir을 절대경로로 정규화해 Session.Meta의 cache_dir에 기록한다.

package runcmd

import (
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/ingest"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestRunRecordsAbsoluteCacheDir(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "session.json")
	if err := quest.New().Save(path); err != nil {
		t.Fatal(err)
	}
	o, cmd := newOptions(t, path)
	o.cacheDir = "warc-cache" // relative; run must normalize to absolute
	stubIngest(t, func(*session.Session, ingest.RunOptions) error { return nil })

	if err := o.run(cmd, nil); err != nil {
		t.Fatalf("run() err = %v", err)
	}
	got, err := quest.Load(path)
	if err != nil {
		t.Fatal(err)
	}
	v, ok := got.GetMeta(session.MetaCacheDir)
	if !ok {
		t.Fatalf("Meta[cache_dir] absent")
	}
	if s, _ := v.(string); !filepath.IsAbs(s) {
		t.Errorf("Meta[cache_dir] = %q, want an absolute path", s)
	}
}
