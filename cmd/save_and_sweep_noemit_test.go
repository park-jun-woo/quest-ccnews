//ff:func feature=cli type=helper control=sequence
//ff:what saveAndSweep가 종단 기사가 없으면 Sweep이 0을 반환해 재저장 없이 세션 저장만 하고 out 파일을 만들지 않는지 검증한다(n==0 분기).

package cmd

import (
	"os"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestSaveAndSweep_NothingToEmit(t *testing.T) {
	sessPath, outPath := setSaveAndSweepPaths(t)
	s := session.New("ua", "cc-news")
	s.Articles = []*session.Article{{
		URL:   "https://example.com/a",
		Host:  "example.com",
		State: session.TODO, // non-terminal ⇒ Sweep returns 0, no re-save
	}}

	if err := saveAndSweep(s); err != nil {
		t.Fatalf("saveAndSweep: %v", err)
	}

	// Session was still saved.
	if _, err := session.Load(sessPath); err != nil {
		t.Fatalf("session not saved: %v", err)
	}
	// Nothing terminal ⇒ no output file created.
	if _, err := os.Stat(outPath); !os.IsNotExist(err) {
		t.Errorf("out file stat err = %v, want not-exist (nothing should be swept)", err)
	}
}
