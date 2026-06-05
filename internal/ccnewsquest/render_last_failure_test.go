//ff:func feature=gate type=helper control=sequence
//ff:what Render 직전 실패 사유 분기 단위테스트. 비어있지 않은 로그 꼬리 Reason이 "직전 실패: …"로 표면화되는지 검증한다(순수, 네트워크 불요).

package ccnewsquest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestRenderLastFailureReason(t *testing.T) {
	// A non-empty log tail Reason is surfaced as 직전 실패.
	it := &quest.Item{
		Key: "https://x/a",
		Log: []quest.Attempt{{Reason: ""}, {Reason: "anchor hallucination"}},
	}
	if err := it.SetPayload(&session.Article{URL: "https://x/a", Host: "x", Lang: "en"}); err != nil {
		t.Fatal(err)
	}
	out, err := Def("ua", "cache").Render(it)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, "직전 실패: anchor hallucination") {
		t.Fatalf("want last failure reason in output:\n%s", out)
	}
}
