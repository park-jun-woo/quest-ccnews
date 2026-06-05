//ff:func feature=gate type=helper control=sequence
//ff:what Render 비-*Article payload 분기 단위테스트. payload가 *Article이 아니면 host/lang은 비어도 URL은 항상 출력하며 렌더가 성공하는지 검증한다(순수, 네트워크 불요).

package ccnewsquest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestRenderNonArticlePayload(t *testing.T) {
	// Non-*Article payload: host/lang stay empty, render still succeeds.
	it := &quest.Item{Key: "https://x/a", Payload: "not an article"}
	out, err := Def("ua", "cache").Render(it)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, "https://x/a") {
		t.Fatalf("URL missing for non-Article payload:\n%s", out)
	}
}
