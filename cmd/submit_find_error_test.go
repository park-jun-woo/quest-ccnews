//ff:func feature=cli type=helper control=sequence level=error
//ff:what runSubmit가 --url이 세션에 없는 기사를 가리킬 때 find 에러를 반환하는지 검증한다.

package cmd

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/spf13/cobra"
)

func TestRunSubmit_FindError(t *testing.T) {
	resetSubmitFlags(t)
	s := session.New("ua", "cc-news") // no articles
	p := filepath.Join(t.TempDir(), "session.json")
	if err := s.Save(p); err != nil {
		t.Fatal(err)
	}
	submitURL = "https://example.com/unknown"
	submitEvent6 = "ev.json"
	sessionPath = p
	cmd := &cobra.Command{}
	cmd.SetOut(new(bytes.Buffer))
	if err := runSubmit(cmd, nil); err == nil {
		t.Fatal("want find error")
	}
}
