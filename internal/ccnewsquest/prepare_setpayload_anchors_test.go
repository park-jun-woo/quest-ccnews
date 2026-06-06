//ff:func feature=gate type=helper control=sequence
//ff:what Prepare가 되쓴 payload의 event6 앵커가 제출 raw와 무손실 라운드트립하는지 검증한다(Phase012). trust PASS 후 DecodePayload로 a2.Event6.Who.Anchors / What.Anchors가 제출 raw의 값과 동일한지 단언.

package ccnewsquest

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestPrepareSetPayloadAnchorsRoundTrip(t *testing.T) {
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

	wantWho := []string{"Rafael Grossi", "IAEA"}
	wantWhat := []string{"framework agreement", "nuclear program"}
	raw := []byte(`{` +
		`"who":{"value":"Rafael Grossi, IAEA","anchors":["Rafael Grossi","IAEA"]},` +
		`"what":{"value":"framework agreement on nuclear program","anchors":["framework agreement","nuclear program"]}` +
		`}`)

	if _, v, err := d.Prepare(quest.New(), it, raw); err != nil || v != nil {
		t.Fatalf("Prepare: v=%+v err=%v", v, err)
	}

	var a2 session.Article
	if err := it.DecodePayload(&a2); err != nil {
		t.Fatalf("DecodePayload: %v", err)
	}
	if a2.Event6 == nil || a2.Event6.Who == nil || a2.Event6.What == nil {
		t.Fatalf("payload event6 missing fields: %+v", a2.Event6)
	}
	if !reflect.DeepEqual(a2.Event6.Who.Anchors, wantWho) {
		t.Fatalf("who.anchors = %v, want %v (lossy round-trip)", a2.Event6.Who.Anchors, wantWho)
	}
	if !reflect.DeepEqual(a2.Event6.What.Anchors, wantWhat) {
		t.Fatalf("what.anchors = %v, want %v (lossy round-trip)", a2.Event6.What.Anchors, wantWhat)
	}
}
