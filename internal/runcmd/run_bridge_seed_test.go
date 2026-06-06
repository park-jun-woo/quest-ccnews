//ff:func feature=ingestion type=command control=sequence
//ff:what run 단위테스트(브리지 시드). 스캔된 기사가 매-WARC Save 콜백(브리지 클로저)으로 reins 세션에 시드되는지 검증한다(Host 빈 문자열로 robots 가드 건너뜀 → 무네트워크).

package runcmd

import (
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/ingest"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestRunBridgesScannedArticle(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "session.json")
	if err := quest.New().Save(path); err != nil {
		t.Fatal(err)
	}
	o, cmd := newOptions(t, path)

	// Simulate a WARC scan appending one article to the scratch, then invoke
	// the per-WARC Save callback (the bridge closure) as the real loop would.
	// Host is left empty so bridge skips the robots guard — no network call.
	stubIngest(t, func(sc *session.Session, opt ingest.RunOptions) error {
		sc.Articles = append(sc.Articles, &session.Article{URL: "https://ex.com/a"})
		return opt.Save()
	})

	if err := o.run(cmd, nil); err != nil {
		t.Fatalf("run() err = %v", err)
	}
	got, err := quest.Load(path)
	if err != nil {
		t.Fatalf("reload: %v", err)
	}
	if len(got.Items) != 1 || got.Items[0].Key != "https://ex.com/a" {
		t.Errorf("items = %+v, want one seeded article", got.Items)
	}
}
