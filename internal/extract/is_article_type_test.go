//ff:func feature=extract type=helper control=iteration dimension=1
//ff:what isArticleType가 기사 문자열/비기사/배열(기사 포함·미포함·비문자열 섞임)/빈배열/nil/숫자를 올바르게 판정하는지 테이블로 검증한다.

package extract

import "testing"

func TestIsArticleType(t *testing.T) {
	cases := []struct {
		name string
		in   any
		want bool
	}{
		{"NewsArticle string", "NewsArticle", true},
		{"Article string", "Article", true},
		{"OpinionNewsArticle", "OpinionNewsArticle", true},
		{"non-article string", "WebPage", false},
		{"array with article", []any{"WebPage", "NewsArticle"}, true},
		{"array without article", []any{"WebPage", "Thing"}, false},
		{"array with non-string", []any{1.0, "Article"}, true},
		{"empty array", []any{}, false},
		{"nil", nil, false},
		{"number", 3.0, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := isArticleType(c.in); got != c.want {
				t.Fatalf("isArticleType(%v) = %v, want %v", c.in, got, c.want)
			}
		})
	}
}
