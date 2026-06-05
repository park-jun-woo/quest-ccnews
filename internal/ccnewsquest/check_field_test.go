//ff:func feature=gate type=helper control=sequence
//ff:what checkField 단위테스트. 유효앵커 전부 substring이면 statusAnchored, 유효앵커 0개(앵커 없음·공백 앵커뿐)면 statusUnanchored, 유효앵커 하나라도 원문에 없으면 statusHallucination(표면형 반환). 공백 앵커는 카운트·매칭 모두 건너뜀.

package ccnewsquest

import "testing"

func TestCheckField(t *testing.T) {
	t.Run("all valid anchors present -> anchored", func(t *testing.T) {
		st, bad := checkField(fld("Alice", "Alice", "Paris"), testSource)
		if st != statusAnchored {
			t.Fatalf("status = %v, want statusAnchored", st)
		}
		if bad != "" {
			t.Fatalf("bad = %q, want empty", bad)
		}
	})
	t.Run("no anchors -> unanchored", func(t *testing.T) {
		st, bad := checkField(fld("Alice"), testSource)
		if st != statusUnanchored {
			t.Fatalf("status = %v, want statusUnanchored", st)
		}
		if bad != "" {
			t.Fatalf("bad = %q, want empty", bad)
		}
	})
	t.Run("whitespace-only anchors skipped -> unanchored", func(t *testing.T) {
		// Both anchors normalize to empty: neither counted nor matched.
		st, _ := checkField(fld("Alice", "  ", ""), testSource)
		if st != statusUnanchored {
			t.Fatalf("status = %v, want statusUnanchored (empty anchors skipped)", st)
		}
	})
	t.Run("hallucinated anchor -> hallucination with surface form", func(t *testing.T) {
		st, bad := checkField(fld("Alice", "Zelda"), testSource)
		if st != statusHallucination {
			t.Fatalf("status = %v, want statusHallucination", st)
		}
		if bad != "Zelda" {
			t.Fatalf("bad = %q, want %q", bad, "Zelda")
		}
	})
	t.Run("empty anchor before a real one is skipped, real one matches", func(t *testing.T) {
		st, _ := checkField(fld("Alice", "   ", "Alice"), testSource)
		if st != statusAnchored {
			t.Fatalf("status = %v, want statusAnchored (blank skipped, real matched)", st)
		}
	})
}
