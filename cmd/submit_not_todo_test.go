//ff:func feature=cli type=helper control=sequence level=error
//ff:what runSubmit가 대상 기사가 TODO가 아닐 때 "TODO가 아님" 에러를 반환하는지 검증한다.

package cmd

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/spf13/cobra"
)

func TestRunSubmit_NotTODO(t *testing.T) {
	resetSubmitFlags(t)
	s := session.New("ua", "cc-news")
	s.Articles = []*session.Article{{URL: "https://example.com/a", State: session.PASS}}
	p := filepath.Join(t.TempDir(), "session.json")
	if err := s.Save(p); err != nil {
		t.Fatal(err)
	}
	submitURL = "https://example.com/a"
	submitEvent6 = "ev.json"
	sessionPath = p
	cmd := &cobra.Command{}
	cmd.SetOut(new(bytes.Buffer))
	if err := runSubmit(cmd, nil); err == nil || !strings.Contains(err.Error(), "TODO가 아님") {
		t.Fatalf("want non-TODO error, got %v", err)
	}
}
