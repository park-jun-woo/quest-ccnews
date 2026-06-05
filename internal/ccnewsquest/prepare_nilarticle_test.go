//ff:func feature=gate type=helper control=sequence
//ff:what Prepare 타입드-nil *Article payload 가드 단위테스트. typed-nil payload가 a == nil 가드에 걸려 에러를 내는지 검증한다(네트워크 불요).

package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestPrepareNilArticlePayload(t *testing.T) {
	// A typed-nil *session.Article payload hits the `a == nil` guard.
	var a *session.Article
	it := &quest.Item{Key: "https://x/a", Payload: a}
	if _, _, err := Def("ua", "cache").Prepare(it, []byte(`{}`)); err == nil {
		t.Fatal("want error for nil *session.Article payload")
	}
}
