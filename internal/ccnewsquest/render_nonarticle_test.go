//ff:func feature=gate type=helper control=sequence
//ff:what Render decode-실패 payload 분기 단위테스트. payload가 Article로 디코드 불가하면 host/lang은 비어도 URL은 항상 출력하며 렌더가 성공하는지 검증한다(순수, 네트워크 불요).

package ccnewsquest

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestRenderNonArticlePayload(t *testing.T) {
	// Payload that does not decode into a session.Article (a bare JSON string):
	// host/lang stay empty (DecodePayload error is swallowed) and render still
	// succeeds with the URL present.
	it := &quest.Item{Key: "https://x/a", Payload: json.RawMessage(`"not an article"`)}
	out, err := Def("ua", "cache").Render(it)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, "https://x/a") {
		t.Fatalf("URL missing for undecodable payload:\n%s", out)
	}
}
