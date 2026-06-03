//ff:func feature=extract type=helper control=sequence
//ff:what firstArticle가 객체 목록에서 첫 Article을 찾아 매핑하고(앞의 비기사는 건너뜀), 기사 없으면 ok=false인지 검증한다.

package extract

import "testing"

func TestFirstArticle(t *testing.T) {
	objs := []map[string]any{
		{"@type": "WebPage", "headline": "skip me"},
		{"@type": "NewsArticle", "headline": "Found"},
		{"@type": "Article", "headline": "Second"},
	}
	f, ok := firstArticle(objs)
	if !ok || f.Title != "Found" {
		t.Fatalf("got %+v ok=%v, want first article 'Found'", f, ok)
	}
	if _, ok := firstArticle([]map[string]any{{"@type": "WebPage"}}); ok {
		t.Fatalf("expected ok=false when no article object")
	}
	if _, ok := firstArticle(nil); ok {
		t.Fatalf("expected ok=false for nil objs")
	}
}
