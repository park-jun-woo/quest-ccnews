//ff:func feature=gate type=helper control=sequence
//ff:what Def().Prepare() 비-*Article payload 가드. payload가 *Article이 아니면 에러를 내는지 검증한다(네트워크 불요).

package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestPrepareNonArticlePayload(t *testing.T) {
	it := &quest.Item{Key: "https://x/a", Payload: "not an article"}
	if _, _, err := Def("ua", "cache").Prepare(it, []byte(`{}`)); err == nil {
		t.Fatal("want error for non-*Article payload")
	}
}
