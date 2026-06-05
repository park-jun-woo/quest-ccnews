//ff:func feature=gate type=helper control=sequence
//ff:what Prepare trust FAIL(short) 경로가 SkipReason을 payload에 보존하는지 검증한다(Phase012). 구조화 없는 본문 재독→추출 SKIP 후 OutSkip verdict, DecodePayload로 a2.SkipReason==기대 사유 단언.

package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/extract"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestPrepareSetPayloadSkip(t *testing.T) {
	cacheDir, file := writeWarcHTML(t, skipHTML)
	d := Def("ua", cacheDir)

	it := &quest.Item{Key: "https://example.com/a"}
	if err := it.SetPayload(&session.Article{
		URL:   "https://example.com/a",
		State: session.TODO,
		WARC:  &session.WARCLoc{File: file, Offset: 0},
	}); err != nil {
		t.Fatal(err)
	}

	raw := []byte(`{"who":{"value":"Reporter","anchors":["Reporter"]}}`)
	_, v, err := d.Prepare(it, raw)
	if err != nil {
		t.Fatalf("Prepare: %v", err)
	}
	if v == nil || v.Outcome != quest.OutSkip {
		t.Fatalf("trust FAIL should short-circuit to OutSkip, got %+v", v)
	}

	var a2 session.Article
	if err := it.DecodePayload(&a2); err != nil {
		t.Fatalf("DecodePayload: %v", err)
	}
	if a2.SkipReason != extract.SkipNoStructured {
		t.Fatalf("payload SkipReason = %q, want %q", a2.SkipReason, extract.SkipNoStructured)
	}
}
