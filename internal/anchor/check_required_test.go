//ff:func feature=anchor type=helper control=sequence
//ff:what checkRequired가 nil/빈 값=FAIL(값 없음), 앵커 환각=FAIL(환각), 앵커 0개=FAIL(검증 불가), 전부 substring=nil(통과)을 반환하는지 분기별로 검증한다.

package anchor

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestCheckRequired(t *testing.T) {
	const norm = "the quick brown fox jumps"

	t.Run("nil field is FAIL", func(t *testing.T) {
		res := checkRequired(namedField{name: "who", f: nil}, norm)
		if res == nil || res.Verdict != FAIL || !strings.Contains(res.Reason, "값 없음") {
			t.Fatalf("res = %+v, want FAIL(값 없음)", res)
		}
	})

	t.Run("empty value is FAIL", func(t *testing.T) {
		f := &session.Field{Value: "", Anchors: []string{"quick"}}
		res := checkRequired(namedField{name: "who", f: f}, norm)
		if res == nil || res.Verdict != FAIL || !strings.Contains(res.Reason, "값 없음") {
			t.Fatalf("res = %+v, want FAIL(값 없음)", res)
		}
	})

	t.Run("no anchors is FAIL (검증 불가)", func(t *testing.T) {
		f := &session.Field{Value: "Alice"} // no anchors
		res := checkRequired(namedField{name: "who", f: f}, norm)
		if res == nil || res.Verdict != FAIL || !strings.Contains(res.Reason, "검증 불가") {
			t.Fatalf("res = %+v, want FAIL(검증 불가)", res)
		}
	})

	t.Run("hallucinated anchor is FAIL (환각)", func(t *testing.T) {
		f := &session.Field{Value: "Alice", Anchors: []string{"lazy dog"}}
		res := checkRequired(namedField{name: "who", f: f}, norm)
		if res == nil || res.Verdict != FAIL || !strings.Contains(res.Reason, "환각") {
			t.Fatalf("res = %+v, want FAIL(환각)", res)
		}
	})

	t.Run("all anchors present passes", func(t *testing.T) {
		f := &session.Field{Value: "Alice", Anchors: []string{"quick", "fox jumps"}}
		res := checkRequired(namedField{name: "who", f: f}, norm)
		if res != nil {
			t.Errorf("res = %+v, want nil (pass)", res)
		}
		if !f.Anchored {
			t.Error("Anchored = false, want true")
		}
	})
}
