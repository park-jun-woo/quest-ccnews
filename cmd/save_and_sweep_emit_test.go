//ff:func feature=cli type=helper control=sequence
//ff:what saveAndSweep가 종단(PASS) 미emit 기사를 out으로 sweep한 뒤 n>0이라 Emitted 플래그를 재저장(두 번째 Save)하는지 검증한다.

package cmd

import (
	"os"
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestSaveAndSweep_EmitsAndReSaves(t *testing.T) {
	sessPath, outPath := setSaveAndSweepPaths(t)
	s := session.New("ua", "cc-news")
	s.Articles = []*session.Article{{
		URL:   "https://example.com/a",
		Host:  "example.com",
		Lang:  "en",
		State: session.PASS, // terminal ⇒ swept to out
	}}

	if err := saveAndSweep(s); err != nil {
		t.Fatalf("saveAndSweep: %v", err)
	}

	// The terminal article was appended to --out.
	data, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("read out: %v", err)
	}
	if !strings.Contains(string(data), `"status":"PASS"`) {
		t.Errorf("out = %q, want a PASS record", string(data))
	}

	// Because n>0, the Emitted flag must have been persisted (the second Save).
	s2, err := session.Load(sessPath)
	if err != nil {
		t.Fatal(err)
	}
	if !s2.Articles[0].Emitted {
		t.Error("persisted Emitted = false, want true after re-save")
	}
}
