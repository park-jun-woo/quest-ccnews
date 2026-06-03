//ff:func feature=anchor type=helper control=sequence
//ff:what checkField가 유효앵커 0개=unanchored, 전부 substring=anchored, 누락=hallucination(원문형 반환), 정규화 후 매칭, Field.Anchored 채움을 검증한다.

package anchor

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestCheckField(t *testing.T) {
	const norm = "the quick brown fox jumps"

	t.Run("no anchors is unanchored", func(t *testing.T) {
		f := &session.Field{Value: "x"}
		st, bad := checkField(f, norm)
		if st != unanchored {
			t.Errorf("status = %v, want unanchored", st)
		}
		if bad != "" {
			t.Errorf("bad = %q, want empty", bad)
		}
		if f.Anchored {
			t.Error("Anchored = true, want false")
		}
	})

	t.Run("all anchors present is anchored", func(t *testing.T) {
		f := &session.Field{Value: "x", Anchors: []string{"quick", "fox jumps"}}
		st, bad := checkField(f, norm)
		if st != anchored {
			t.Errorf("status = %v, want anchored", st)
		}
		if bad != "" {
			t.Errorf("bad = %q, want empty", bad)
		}
		if !f.Anchored {
			t.Error("Anchored = false, want true")
		}
	})

	t.Run("missing anchor is hallucination", func(t *testing.T) {
		f := &session.Field{Value: "x", Anchors: []string{"quick", "lazy dog"}}
		st, bad := checkField(f, norm)
		if st != hallucination {
			t.Errorf("status = %v, want hallucination", st)
		}
		if bad != "lazy dog" {
			t.Errorf("bad = %q, want %q", bad, "lazy dog")
		}
		if f.Anchored {
			t.Error("Anchored = true, want false")
		}
	})

	t.Run("anchor normalized before match", func(t *testing.T) {
		f := &session.Field{Value: "x", Anchors: []string{"  quick   brown  "}}
		st, _ := checkField(f, norm)
		if st != anchored {
			t.Errorf("status = %v, want anchored (normalized anchor should match)", st)
		}
		if !f.Anchored {
			t.Error("Anchored = false, want true")
		}
	})
}
