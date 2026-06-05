//ff:func feature=gate type=helper control=sequence
//ff:what Render 빈 사유 분기 단위테스트. 로그는 있으나 꼬리 Reason이 비면 "직전 실패:" 줄을 출력하지 않는지 검증한다(순수, 네트워크 불요).

package ccnewsquest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestRenderEmptyReasonNotShown(t *testing.T) {
	// Log present but tail Reason empty -> no 직전 실패 line.
	it := &quest.Item{
		Key: "https://x/a",
		Log: []quest.Attempt{{Reason: ""}},
	}
	if err := it.SetPayload(&session.Article{URL: "https://x/a"}); err != nil {
		t.Fatal(err)
	}
	out, err := Def("ua", "cache").Render(it)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(out, "직전 실패:") {
		t.Fatalf("empty tail reason must not print 직전 실패 line:\n%s", out)
	}
}
