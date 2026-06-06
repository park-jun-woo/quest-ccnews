//ff:func feature=gate type=helper control=sequence
//ff:what Def().Prepare() decode-실패 payload 가드. payload가 Article로 디코드 불가하면 에러를 내는지 검증한다(네트워크 불요).

package ccnewsquest

import (
	"encoding/json"
	"testing"

	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestPrepareNonArticlePayload(t *testing.T) {
	// A payload that cannot decode into session.Article (a bare JSON string)
	// makes DecodePayload fail, which Prepare surfaces as an error.
	it := &quest.Item{Key: "https://x/a", Payload: json.RawMessage(`"not an article"`)}
	if _, _, err := Def("ua", "cache").Prepare(quest.New(), it, []byte(`{}`)); err == nil {
		t.Fatal("want error for payload that does not decode into *session.Article")
	}
}
