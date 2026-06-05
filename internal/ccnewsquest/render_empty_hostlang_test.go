//ff:func feature=gate type=helper control=sequence
//ff:what Render 빈 host/lang 분기 단위테스트. host·lang이 비면 해당 줄을 생략하고 URL·submit 사용법은 항상 출력하는지 검증한다(순수, 네트워크 불요).

package ccnewsquest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestRenderEmptyHostLang(t *testing.T) {
	// Empty host and lang: those lines must be omitted, URL/submit still present.
	it := &quest.Item{Key: "https://x/a", Payload: &session.Article{URL: "https://x/a"}}
	out, err := Def("ua", "cache").Render(it)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(out, "호스트:") {
		t.Fatalf("host line should be omitted when host empty:\n%s", out)
	}
	if strings.Contains(out, "언어:") {
		t.Fatalf("lang line should be omitted when lang empty:\n%s", out)
	}
	if !strings.Contains(out, "https://x/a") || !strings.Contains(out, "submit --key") {
		t.Fatalf("URL/submit usage missing:\n%s", out)
	}
}
