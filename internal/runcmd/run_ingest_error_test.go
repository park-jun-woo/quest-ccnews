//ff:func feature=ingestion type=command control=sequence level=error
//ff:what run 단위테스트(ingest 루프 에러). ingestRun이 에러를 반환하면 run이 그 에러를 그대로 전파한다(errors.Is).

package runcmd

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/ingest"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestRunIngestLoopErrorPropagates(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "session.json")
	if err := quest.New().Save(path); err != nil {
		t.Fatal(err)
	}
	o, cmd := newOptions(t, path)

	wantErr := errors.New("boom")
	stubIngest(t, func(*session.Session, ingest.RunOptions) error { return wantErr })

	if err := o.run(cmd, nil); !errors.Is(err, wantErr) {
		t.Errorf("run() err = %v, want boom", err)
	}
}
