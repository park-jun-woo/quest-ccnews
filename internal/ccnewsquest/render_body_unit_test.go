//ff:func feature=gate type=helper control=sequence
//ff:what renderBody 단위테스트(분기 직접 커버). ① PASS 본문 → "원문 본문" 섹션 + 본문 텍스트를 b에 기록. ② 신뢰 게이트 SKIP(readArticleBody ok=false) → 아무것도 기록 안 함(best-effort, 본문 생략). 두 경우 모두 Meta 미변경(read-only).

package ccnewsquest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestRenderBodyBranches(t *testing.T) {
	t.Run("pass body is appended", func(t *testing.T) {
		cacheDir, file := writeWarcHTML(t, passHTML)
		d := Def("ua", cacheDir).(ccnewsDef)
		it := &quest.Item{Key: "https://example.com/a"}
		if err := it.SetPayload(&session.Article{
			URL: "https://example.com/a", State: session.TODO,
			WARC: &session.WARCLoc{File: file, Offset: 0},
		}); err != nil {
			t.Fatal(err)
		}
		var b strings.Builder
		d.renderBody(quest.New(), it, &b)
		if !strings.Contains(b.String(), passBody) {
			t.Fatalf("renderBody omitted body text:\n%s", b.String())
		}
	})

	t.Run("trust SKIP appends nothing", func(t *testing.T) {
		cacheDir, file := writeWarcHTML(t, skipHTML)
		d := Def("ua", cacheDir).(ccnewsDef)
		it := &quest.Item{Key: "https://example.com/a"}
		if err := it.SetPayload(&session.Article{
			URL: "https://example.com/a", State: session.TODO,
			WARC: &session.WARCLoc{File: file, Offset: 0},
		}); err != nil {
			t.Fatal(err)
		}
		var b strings.Builder
		d.renderBody(quest.New(), it, &b)
		if b.String() != "" {
			t.Fatalf("renderBody wrote on trust SKIP: %q", b.String())
		}
	})
}
