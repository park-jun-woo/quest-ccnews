//ff:func feature=ingestion type=command control=sequence
//ff:what run 단위테스트. 세션 파일 부재면 빈 세션으로 시작·실행·최종 flush로 저장, 손상된 세션은 Load 에러 전파, 잘못된 ingestion Meta는 restoreScratch 에러 전파, ingest 루프 에러 전파, 스캔된 기사가 매-WARC Save 콜백(브리지)으로 reins 세션에 시드되는지 검증한다(ingestRun 스텁으로 무네트워크).

package runcmd

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/ingest"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestRun(t *testing.T) {
	t.Run("missing session file → starts fresh, runs, no error", func(t *testing.T) {
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
	})

	t.Run("malformed session file → Load error propagates", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "session.json")
		if err := os.WriteFile(path, []byte("{not json"), 0o644); err != nil {
			t.Fatal(err)
		}
		o, cmd := newOptions(t, path)

		if err := o.run(cmd, nil); err == nil {
			t.Errorf("run() err = nil, want unmarshal error")
		}
	})

	t.Run("bad ingestion Meta → restoreScratch error propagates", func(t *testing.T) {
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
	})

	t.Run("ingest loop error propagates", func(t *testing.T) {
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
	})

	t.Run("scanned article is bridged into the reins session", func(t *testing.T) {
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
	})
}
