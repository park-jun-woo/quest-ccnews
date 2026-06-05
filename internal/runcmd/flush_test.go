//ff:func feature=ingestion type=command control=sequence
//ff:what flush 단위테스트. 기사를 시드·저장하고 seed 카운트를 보고하며 세션 파일에 Item이 영속되는지, 새 기사가 없으면 조용히 저장(seed 줄 없음)하는지, Save 실패가 전파되는지 검증한다.

package runcmd

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestFlush(t *testing.T) {
	t.Run("seeds articles, saves, and reports the count", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "session.json")
		o := flushOptions(t, path)

		scratch := session.New("ua", "cc-news")
		scratch.Articles = []*session.Article{{URL: "https://ex.com/a"}}
		s := quest.New()
		var buf bytes.Buffer

		if err := o.flush(scratch, s, nil, "2026-06-05T00:00:00Z", &buf); err != nil {
			t.Fatalf("flush() err = %v", err)
		}
		if !strings.Contains(buf.String(), "seed: +1 items") {
			t.Errorf("output = %q, want seed report", buf.String())
		}
		// Persisted session has the seeded item.
		got, err := quest.Load(path)
		if err != nil {
			t.Fatalf("reload: %v", err)
		}
		if len(got.Items) != 1 {
			t.Errorf("items = %d, want 1", len(got.Items))
		}
	})

	t.Run("no new articles → saves silently (no seed line)", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "session.json")
		o := flushOptions(t, path)

		scratch := session.New("ua", "cc-news") // no Articles
		s := quest.New()
		var buf bytes.Buffer

		if err := o.flush(scratch, s, nil, "now", &buf); err != nil {
			t.Fatalf("flush() err = %v", err)
		}
		if buf.Len() != 0 {
			t.Errorf("output = %q, want empty (n==0)", buf.String())
		}
	})

	t.Run("Save error propagates", func(t *testing.T) {
		// A session path under a non-existent directory makes WriteFile fail.
		o := flushOptions(t, filepath.Join(t.TempDir(), "no-such-dir", "session.json"))
		scratch := session.New("ua", "cc-news")
		s := quest.New()
		var buf bytes.Buffer

		if err := o.flush(scratch, s, nil, "now", &buf); err == nil {
			t.Errorf("flush() err = nil, want Save error")
		}
	})
}
