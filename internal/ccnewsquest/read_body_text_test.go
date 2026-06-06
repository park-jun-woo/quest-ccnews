//ff:func feature=gate type=helper control=sequence level=error
//ff:what readArticleBody 단위 테스트. WARC 재독+extract.Apply 공유 헬퍼의 분기를 직접 커버한다. ① PASS: 본문 텍스트(=게이트 Source)와 ok=true 반환, a를 Extracted/Lang으로 제자리 변이. ② ReadBody 실패(없는 파일): err 반환·ok=false·빈 본문. ③ 신뢰 게이트 SKIP(구조화 데이터 없음): ok=false·State=SKIPPED·SkipReason 설정·err 없음.
package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestReadArticleBodyPass(t *testing.T) {
	cacheDir, file := writeWarcHTML(t, passHTML)
	a := &session.Article{
		URL:   "https://example.com/a",
		State: session.TODO,
		WARC:  &session.WARCLoc{File: file, Offset: 0},
	}

	body, ok, err := readArticleBody("ua", cacheDir, a)
	if err != nil {
		t.Fatalf("readArticleBody: %v", err)
	}
	if !ok {
		t.Fatalf("ok=false, want true (passHTML clears the trust gate)")
	}
	if body != passBody {
		t.Fatalf("body = %q, want passBody (byte-identical anchor target)", body)
	}
	if a.Extracted == nil {
		t.Errorf("Extracted not filled on PASS")
	}
	if a.Lang != "en" {
		t.Errorf("Lang = %q, want en (filled from extraction)", a.Lang)
	}
}
