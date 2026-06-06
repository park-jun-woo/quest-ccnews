//ff:func feature=gate type=helper control=sequence
//ff:what Phase013 C 검증 게이트. ① Render(next) 출력에 앵커 대상 본문 텍스트가 포함된다. ② 동일 기사·동일 cacheDir에 대해 Render가 보여준 본문 == Prepare의 Context.Source(바이트 동일, 공유 헬퍼 readArticleBody). ③ next 출력 본문에서 고른 표면형 앵커가 submit(Prepare→checkField)에서 그대로 anchored 통과(라운드트립).
package ccnewsquest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestRenderIncludesBodyMatchingPrepareSource(t *testing.T) {
	cacheDir, file := writeWarcHTML(t, passHTML)
	d := Def("ua", cacheDir)

	newItem := func() *quest.Item {
		it := &quest.Item{Key: "https://example.com/a"}
		if err := it.SetPayload(&session.Article{
			URL:   "https://example.com/a",
			State: session.TODO,
			WARC:  &session.WARCLoc{File: file, Offset: 0},
		}); err != nil {
			t.Fatal(err)
		}
		return it
	}

	// ① Render output includes the anchor-target body.
	out, err := d.Render(quest.New(), newItem())
	if err != nil {
		t.Fatalf("Render: %v", err)
	}
	if !strings.Contains(out, passBody) {
		t.Fatalf("Render output missing body text:\n%s", out)
	}

	// ② The body shown by next equals Prepare's Context.Source byte-for-byte.
	ctx, v, err := d.Prepare(quest.New(), newItem(), []byte(`{"who":{"value":"x","anchors":["article body"]}}`))
	if err != nil || v != nil {
		t.Fatalf("Prepare: v=%+v err=%v", v, err)
	}
	if !strings.Contains(out, ctx.Source) {
		t.Fatalf("Render body does not contain Prepare Source byte-identically")
	}
	// And the body Render shows is exactly Prepare's Source (same helper).
	if ctx.Source != passBody {
		t.Fatalf("Prepare Source = %q, want passBody", ctx.Source)
	}

	// ③ A surface anchor visible in the next body passes submit (anchored==true).
	it := newItem()
	if _, v, err := d.Prepare(quest.New(), it, []byte(`{"who":{"value":"x","anchors":["article body"]}}`)); err != nil || v != nil {
		t.Fatalf("Prepare round-trip: v=%+v err=%v", v, err)
	}
	var a2 session.Article
	if err := it.DecodePayload(&a2); err != nil {
		t.Fatal(err)
	}
	if a2.Event6 == nil || a2.Event6.Who == nil || !a2.Event6.Who.Anchored {
		t.Fatalf("anchor picked from next body did not pass submit: %+v", a2.Event6)
	}
}
