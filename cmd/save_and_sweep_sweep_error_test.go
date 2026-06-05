//ff:func feature=cli type=helper control=sequence level=error
//ff:what saveAndSweep가 세션 저장 후 종단 기사를 out으로 sweep할 때, out 경로의 부모가 파일이라 append(MkdirAll)가 실패하면 그 에러를 반환하는지 검증한다.

package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestSaveAndSweep_SweepError(t *testing.T) {
	dir := t.TempDir()
	prevSess, prevOut := sessionPath, submitOut
	sessionPath = filepath.Join(dir, "session.json")
	t.Cleanup(func() { sessionPath, submitOut = prevSess, prevOut })

	// A regular file used as a directory component of --out makes Append's
	// MkdirAll (inside Sweep) fail.
	blocker := filepath.Join(dir, "blocker")
	if err := os.WriteFile(blocker, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	submitOut = filepath.Join(blocker, "sub", "out.jsonl")

	s := session.New("ua", "cc-news")
	s.Articles = []*session.Article{{
		URL:   "https://example.com/a",
		Host:  "example.com",
		State: session.PASS, // terminal ⇒ Sweep tries to append ⇒ error
	}}

	if err := saveAndSweep(s); err == nil {
		t.Fatal("want sweep append error")
	}
}
