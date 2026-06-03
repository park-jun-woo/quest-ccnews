//ff:func feature=extract type=helper control=sequence
//ff:what anyArticleType가 기사 타입 문자열을 포함한 배열은 true, 미포함·비문자열만·빈 배열은 false인지 검증한다.

package extract

import "testing"

func TestAnyArticleType(t *testing.T) {
	if !anyArticleType([]any{"WebPage", "NewsArticle"}) {
		t.Fatalf("array containing NewsArticle should be true")
	}
	if anyArticleType([]any{"WebPage", "Thing"}) {
		t.Fatalf("array without article type should be false")
	}
	if !anyArticleType([]any{1.0, "Article"}) {
		t.Fatalf("non-string element skipped, Article still found")
	}
	if anyArticleType([]any{}) {
		t.Fatalf("empty array should be false")
	}
}
