//ff:func feature=gate type=helper control=sequence
//ff:what Phase013 D 검증 게이트. Prepare가 SetPayload 전에 각 present 필드의 anchored를 게이트 동일 함수(checkField)로 채우는지 검증한다. ① 앵커가 본문 substring인 필드(PASS) anchored==true. ② 유효앵커 0개 선택필드 anchored==false. ③ 각 필드 payload anchored가 같은 Source·같은 checkField 결과와 정확히 일치(verdict 모순 0).
package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestPrepareFillsAnchored(t *testing.T) {
	cacheDir, file := writeWarcHTML(t, passHTML)
	d := Def("ua", cacheDir)

	it := &quest.Item{Key: "https://example.com/a"}
	if err := it.SetPayload(&session.Article{
		URL:   "https://example.com/a",
		State: session.TODO,
		WARC:  &session.WARCLoc{File: file, Offset: 0},
	}); err != nil {
		t.Fatal(err)
	}

	// who/what anchors are substrings of passBody (anchored); when has a single
	// valid anchor absent from the body would FAIL the field — instead use a field
	// with zero valid anchors (whitespace only) to exercise the unanchored→false path.
	raw := []byte(`{` +
		`"who":{"value":"body","anchors":["article body"]},` +
		`"what":{"value":"length","anchors":["minimum length"]},` +
		`"when":{"value":"x","anchors":["   "]}` +
		`}`)

	ctx, v, err := d.Prepare(quest.New(), it, raw)
	if err != nil {
		t.Fatalf("Prepare: %v", err)
	}
	if v != nil {
		t.Fatalf("trust PASS should not short-circuit, got %+v", v)
	}

	var a2 session.Article
	if err := it.DecodePayload(&a2); err != nil {
		t.Fatalf("DecodePayload: %v", err)
	}
	ev := a2.Event6
	if ev == nil || ev.Who == nil || ev.What == nil || ev.When == nil {
		t.Fatalf("payload event6 missing fields: %+v", ev)
	}

	// ① anchored fields → true
	if !ev.Who.Anchored {
		t.Errorf("who.anchored = false, want true (anchor is a source substring)")
	}
	if !ev.What.Anchored {
		t.Errorf("what.anchored = false, want true")
	}
	// ② zero valid anchors → false
	if ev.When.Anchored {
		t.Errorf("when.anchored = true, want false (no valid anchor)")
	}

	// ③ each payload anchored equals the gate's checkField over the SAME Source
	// (zero contradictions: same function, same field order).
	source := ctx.Source
	assertAnchoredMatchesCheckField(t, "who", ev.Who, source)
	assertAnchoredMatchesCheckField(t, "what", ev.What, source)
	assertAnchoredMatchesCheckField(t, "when", ev.When, source)
}
