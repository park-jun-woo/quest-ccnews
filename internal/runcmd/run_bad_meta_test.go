//ff:func feature=ingestion type=command control=sequence level=error
//ff:what run 단위테스트(잘못된 ingestion Meta). ingestion Meta가 객체가 아니면 restoreScratch 에러가 run에서 전파된다.

package runcmd

import (
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestRunBadIngestionMetaError(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "session.json")
	s := quest.New()
	s.SetMeta(metaIngestion, "not-an-object")
	if err := s.Save(path); err != nil {
		t.Fatal(err)
	}
	o, cmd := newOptions(t, path)

	if err := o.run(cmd, nil); err == nil {
		t.Errorf("run() err = nil, want restoreScratch error")
	}
}
