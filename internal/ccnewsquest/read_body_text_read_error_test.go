//ff:func feature=gate type=helper control=sequence level=error
//ff:what readArticleBody 단위 테스트(ReadBody 실패). 없는 WARC 파일 → err 반환·ok=false·빈 본문.
package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestReadArticleBodyReadError(t *testing.T) {
	a := &session.Article{
		URL:   "https://example.com/a",
		State: session.TODO,
		WARC:  &session.WARCLoc{File: "does-not-exist.warc", Offset: 0},
	}
	body, ok, err := readArticleBody("ua", t.TempDir(), a)
	if err == nil {
		t.Fatal("want error when ReadBody cannot open the WARC file")
	}
	if ok || body != "" {
		t.Fatalf("got (%q,%v) on error, want (\"\",false)", body, ok)
	}
}
