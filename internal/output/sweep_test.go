//ff:func feature=output type=helper control=sequence level=error
//ff:what Sweep가 nil 세션·종단 미emit 기사만 append·Emitted 표시·emit-once 재호출·append 오류 시 카운트와 미표시 보존을 검증한다.

package output

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestSweep(t *testing.T) {
	t.Run("nil session", func(t *testing.T) {
		n, err := Sweep(nil, filepath.Join(t.TempDir(), "out.jsonl"))
		if err != nil || n != 0 {
			t.Errorf("Sweep(nil) = (%d, %v), want (0, nil)", n, err)
		}
	})

	t.Run("emits terminal only and marks", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "out.jsonl")

		s := &session.Session{
			Hosts: map[string]*session.Host{
				"e.com": {MediaName: "M", Robots: &session.Robots{CrawlAllowed: true}},
			},
			Articles: []*session.Article{
				{URL: "p", Host: "e.com", State: session.PASS},
				{URL: "r", Host: "e.com", State: session.REVIEW},
				{URL: "b", Host: "e.com", State: session.BLOCKED, SkipReason: "x"},
				{URL: "s", Host: "e.com", State: session.SKIPPED, SkipReason: "y"},
				{URL: "todo", Host: "e.com", State: session.TODO},                   // skipped: non-terminal
				{URL: "done", Host: "e.com", State: session.DONE},                   // emitted: DONE is audit
				{URL: "already", Host: "e.com", State: session.PASS, Emitted: true}, // skipped: emitted
			},
		}

		n, err := Sweep(s, path)
		if err != nil {
			t.Fatalf("Sweep() error = %v", err)
		}
		if n != 5 {
			t.Errorf("Sweep() wrote %d, want 5", n)
		}
		if got := countLines(t, path); got != 5 {
			t.Errorf("file has %d lines, want 5", got)
		}

		for _, a := range s.Articles {
			wantEmitted := a.URL == "p" || a.URL == "r" || a.URL == "b" || a.URL == "s" || a.URL == "done" || a.URL == "already"
			if a.Emitted != wantEmitted {
				t.Errorf("article %q Emitted = %v, want %v", a.URL, a.Emitted, wantEmitted)
			}
		}
	})

	t.Run("emit once", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "out.jsonl")

		s := &session.Session{
			Hosts: map[string]*session.Host{},
			Articles: []*session.Article{
				{URL: "p", Host: "e.com", State: session.PASS},
				{URL: "q", Host: "e.com", State: session.REVIEW},
			},
		}

		n1, err := Sweep(s, path)
		if err != nil || n1 != 2 {
			t.Fatalf("first Sweep() = (%d, %v), want (2, nil)", n1, err)
		}

		n2, err := Sweep(s, path)
		if err != nil || n2 != 0 {
			t.Errorf("second Sweep() = (%d, %v), want (0, nil)", n2, err)
		}
		if got := countLines(t, path); got != 2 {
			t.Errorf("file has %d lines after re-sweep, want 2 (no duplicates)", got)
		}
	})

	t.Run("append error returns count", func(t *testing.T) {
		dir := t.TempDir()
		blocker := filepath.Join(dir, "blocker")
		if err := os.WriteFile(blocker, []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		path := filepath.Join(blocker, "child", "out.jsonl")

		s := &session.Session{
			Hosts:    map[string]*session.Host{},
			Articles: []*session.Article{{URL: "p", Host: "e.com", State: session.PASS}},
		}
		n, err := Sweep(s, path)
		if err == nil {
			t.Fatal("Sweep() error = nil, want append error")
		}
		if n != 0 {
			t.Errorf("Sweep() count = %d, want 0", n)
		}
		if s.Articles[0].Emitted {
			t.Errorf("article Emitted = true after failed append, want false")
		}
	})
}
