//ff:func feature=gate type=helper control=sequence
//ff:what Def().Prepare() 잘못된 JSON 가드. 디코드 불가능한 입력에서 에러를 내는지 검증한다(네트워크 불요).

package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestPrepareBadJSON(t *testing.T) {
	it := &quest.Item{Key: "https://x/a"}
	if err := it.SetPayload(&session.Article{URL: "https://x/a"}); err != nil {
		t.Fatal(err)
	}
	if _, _, err := Def("ua", "cache").Prepare(it, []byte("{not json")); err == nil {
		t.Fatal("want JSON decode error")
	}
}
