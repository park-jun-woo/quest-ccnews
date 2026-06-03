//ff:func feature=extract type=helper control=sequence
//ff:what mapArticle가 필드 없는 객체에 대해 zero Fields를 돌려주는지 검증한다.

package extract

import "testing"

func TestMapArticleMissingFields(t *testing.T) {
	got := mapArticle(map[string]any{"@type": "Article"})
	if got != (Fields{}) {
		t.Fatalf("mapArticle of empty = %+v, want zero Fields", got)
	}
}
