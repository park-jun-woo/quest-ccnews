//ff:func feature=anchor type=helper control=sequence
//ff:what checkField가 빈/공백 앵커를 무효로 무시함을 검증한다(Phase009 L0): [""]·["  "]는 unanchored, 빈+진짜는 anchored, 빈+없는앵커는 hallucination.

package anchor

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestCheckField_EmptyAnchor(t *testing.T) {
	const norm = "the quick brown fox jumps"

	t.Run("empty anchor only is unanchored", func(t *testing.T) {
		f := &session.Field{Value: "x", Anchors: []string{""}}
		st, bad := checkField(f, norm)
		if st != unanchored {
			t.Errorf("status = %v, want unanchored ([\"\"] must not anchor — Phase009 L0)", st)
		}
		if bad != "" {
			t.Errorf("bad = %q, want empty", bad)
		}
		if f.Anchored {
			t.Error("Anchored = true, want false ([\"\"] cheese vector)")
		}
	})

	t.Run("whitespace anchors only is unanchored", func(t *testing.T) {
		f := &session.Field{Value: "x", Anchors: []string{"  ", "\t\n"}}
		st, _ := checkField(f, norm)
		if st != unanchored {
			t.Errorf("status = %v, want unanchored (whitespace-only anchors are invalid)", st)
		}
		if f.Anchored {
			t.Error("Anchored = true, want false")
		}
	})

	t.Run("empty plus real anchor is anchored", func(t *testing.T) {
		f := &session.Field{Value: "x", Anchors: []string{"", "quick"}}
		st, _ := checkField(f, norm)
		if st != anchored {
			t.Errorf("status = %v, want anchored (one valid real anchor present)", st)
		}
		if !f.Anchored {
			t.Error("Anchored = false, want true")
		}
	})

	t.Run("empty plus hallucinated anchor is hallucination", func(t *testing.T) {
		f := &session.Field{Value: "x", Anchors: []string{"", "lazy dog"}}
		st, bad := checkField(f, norm)
		if st != hallucination {
			t.Errorf("status = %v, want hallucination (real anchor absent overrides empties)", st)
		}
		if bad != "lazy dog" {
			t.Errorf("bad = %q, want %q", bad, "lazy dog")
		}
		if f.Anchored {
			t.Error("Anchored = true, want false")
		}
	})
}
