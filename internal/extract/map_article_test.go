//ff:func feature=extract type=helper control=sequence
//ff:what mapArticle가 NewsArticle 객체의 headline/author/datePublished/publisher/inLanguage/articleBody를 Fields로 올바르게 매핑하는지 검증한다.

package extract

import "testing"

func TestMapArticle(t *testing.T) {
	o := map[string]any{
		"@type":         "NewsArticle",
		"headline":      "  Big News  ",
		"author":        map[string]any{"name": "Jane"},
		"datePublished": "2026-01-02T03:04:05Z",
		"publisher":     map[string]any{"name": "Acme"},
		"inLanguage":    "en",
		"articleBody":   "body text",
	}
	got := mapArticle(o)
	want := Fields{
		Title:       "Big News",
		Author:      "Jane",
		PublishedAt: "2026-01-02T03:04:05Z",
		MediaName:   "Acme",
		Lang:        "en",
		Body:        "body text",
	}
	if got != want {
		t.Fatalf("mapArticle = %+v, want %+v", got, want)
	}
}
