//ff:func feature=gate type=helper control=sequence
//ff:what Prepare trust PASS 경로가 enrich한 Article을 payload에 되쓰는지 검증한다(Phase012). 재독→추출 PASS 후 DecodePayload로 a2.Event6!=nil 그리고 a2.Extracted(Title/Source 비어있지 않음) 단언.

package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestPrepareSetPayloadPass(t *testing.T) {
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

	raw := []byte(`{"who":{"value":"Reporter","anchors":["Reporter"]}}`)
	_, v, err := d.Prepare(quest.New(), it, raw)
	if err != nil {
		t.Fatalf("Prepare: %v", err)
	}
	if v != nil {
		t.Fatalf("trust PASS should not short-circuit, got verdict %+v", v)
	}

	var a2 session.Article
	if err := it.DecodePayload(&a2); err != nil {
		t.Fatalf("DecodePayload: %v", err)
	}
	if a2.Event6 == nil {
		t.Fatal("payload Event6 nil: Prepare did not write event6 back")
	}
	if a2.Extracted == nil {
		t.Fatal("payload Extracted nil: Prepare did not write extraction back")
	}
	if a2.Extracted.Title == "" {
		t.Errorf("Extracted.Title empty")
	}
	if a2.Extracted.Source == "" {
		t.Errorf("Extracted.Source empty")
	}
}
