//ff:func feature=extract type=helper control=sequence
//ff:what hasArticle가 Title이 있으면 true, 비어있으면 false를 돌려주는지 검증한다.

package extract

import "testing"

func TestHasArticle(t *testing.T) {
	if !(Fields{Title: "x"}).hasArticle() {
		t.Fatalf("title present should be true")
	}
	if (Fields{}).hasArticle() {
		t.Fatalf("empty title should be false")
	}
}
