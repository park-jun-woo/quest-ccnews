//ff:func feature=gate type=helper control=sequence level=error
//ff:what readArticleBody 단위 테스트(신뢰 게이트 SKIP). 구조화 데이터 없는 본문 → ok=false·State=SKIPPED·SkipReason 설정·err 없음.
package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestReadArticleBodyTrustSkip(t *testing.T) {
	cacheDir, file := writeWarcHTML(t, skipHTML)
	a := &session.Article{
		URL:   "https://example.com/a",
		State: session.TODO,
		WARC:  &session.WARCLoc{File: file, Offset: 0},
	}
	body, ok, err := readArticleBody("ua", cacheDir, a)
	if err != nil {
		t.Fatalf("readArticleBody: %v (a trust SKIP is not an error)", err)
	}
	if ok || body != "" {
		t.Fatalf("got (%q,%v), want (\"\",false) on trust SKIP", body, ok)
	}
	if a.State != session.SKIPPED || a.SkipReason == "" {
		t.Errorf("a State=%q reason=%q, want SKIPPED + a skip reason", a.State, a.SkipReason)
	}
}
