//ff:func feature=cli type=helper control=sequence
//ff:what runNext가 TODO 기사가 없을 때 "처리할 TODO 기사가 없습니다" 안내를 출력하고 에러 없이 끝나는지 검증한다.

package cmd

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/spf13/cobra"
)

func TestRunNext_NoTODO(t *testing.T) {
	s := session.New("ua", "cc-news") // no articles
	p := filepath.Join(t.TempDir(), "session.json")
	if err := s.Save(p); err != nil {
		t.Fatal(err)
	}
	sessionPath = p
	t.Cleanup(func() { sessionPath = "session.json" })

	cmd := &cobra.Command{}
	var out bytes.Buffer
	cmd.SetOut(&out)
	if err := runNext(cmd, nil); err != nil {
		t.Fatalf("runNext: %v", err)
	}
	if !strings.Contains(out.String(), "처리할 TODO 기사가 없습니다") {
		t.Errorf("output = %q", out.String())
	}
}
