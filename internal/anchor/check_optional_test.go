//ff:func feature=anchor type=helper control=sequence
//ff:what checkOptional이 nil 무시, 앵커 환각=FAIL Result, 앵커 0개=(nil,true) REVIEW 후보, 전부 substring=(nil,false)를 반환하는지 분기별로 검증한다.

package anchor

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestCheckOptional(t *testing.T) {
	const norm = "the quick brown fox jumps"

	t.Run("nil field is ignored", func(t *testing.T) {
		res, unanchored := checkOptional(namedField{name: "where", f: nil}, norm)
		if res != nil {
			t.Errorf("res = %+v, want nil", res)
		}
		if unanchored {
			t.Error("unanchored = true, want false for nil field")
		}
	})

	t.Run("hallucinated anchor is FAIL", func(t *testing.T) {
		f := &session.Field{Value: "x", Anchors: []string{"lazy dog"}}
		res, unanchored := checkOptional(namedField{name: "where", f: f}, norm)
		if res == nil || res.Verdict != FAIL {
			t.Fatalf("res = %+v, want FAIL", res)
		}
		if !strings.Contains(res.Reason, "where") || !strings.Contains(res.Reason, "환각") {
			t.Errorf("Reason %q should name field and hallucination", res.Reason)
		}
		if unanchored {
			t.Error("unanchored = true, want false on hallucination")
		}
	})

	t.Run("no anchors is REVIEW candidate", func(t *testing.T) {
		f := &session.Field{Value: "x"} // no anchors
		res, unanchored := checkOptional(namedField{name: "why", f: f}, norm)
		if res != nil {
			t.Errorf("res = %+v, want nil", res)
		}
		if !unanchored {
			t.Error("unanchored = false, want true for anchorless present field")
		}
	})

	t.Run("all anchors present passes", func(t *testing.T) {
		f := &session.Field{Value: "x", Anchors: []string{"quick", "fox jumps"}}
		res, unanchored := checkOptional(namedField{name: "how", f: f}, norm)
		if res != nil {
			t.Errorf("res = %+v, want nil", res)
		}
		if unanchored {
			t.Error("unanchored = true, want false when anchored")
		}
		if !f.Anchored {
			t.Error("Anchored = false, want true")
		}
	})
}
