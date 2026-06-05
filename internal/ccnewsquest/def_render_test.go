//ff:func feature=gate type=helper control=sequence
//ff:what Def().Render() 스모크. 아이템 출력에 URL과 submit 사용법(submit --key)이 들어가는지 검증한다.

package ccnewsquest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestRender(t *testing.T) {
	it := &quest.Item{Key: "https://x/a", Tries: 1}
	if err := it.SetPayload(&session.Article{URL: "https://x/a", Host: "x", Lang: "en"}); err != nil {
		t.Fatal(err)
	}
	out, err := Def("ua", "cache").Render(it)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, "https://x/a") || !strings.Contains(out, "submit --key") {
		t.Fatalf("render missing URL or submit usage:\n%s", out)
	}
}
