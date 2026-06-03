//ff:func feature=cli type=helper control=sequence
//ff:what 테스트 헬퍼. warcFile offset 0을 가리키는 TODO 기사 1건의 세션을 임시 session.json에 쓰고 그 경로를 돌려준다.

package cmd

import (
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// writeSessionWith builds a session with one TODO article pointing at warcFile
// offset 0, writes it to a temp session.json, and returns its path.
func writeSessionWith(t *testing.T, warcFile string) string {
	t.Helper()
	s := session.New("ua", "cc-news")
	s.Articles = []*session.Article{{
		URL:   "https://example.com/a",
		Host:  "example.com",
		Lang:  "en",
		State: session.TODO,
		WARC:  &session.WARCLoc{File: warcFile, Offset: 0},
	}}
	p := filepath.Join(t.TempDir(), "session.json")
	if err := s.Save(p); err != nil {
		t.Fatal(err)
	}
	return p
}
