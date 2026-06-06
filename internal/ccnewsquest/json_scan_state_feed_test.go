//ff:func feature=event6 type=helper control=sequence
//ff:what (*jsonScanState).feed 분기 테스트. '{' 첫위치·depth 증가, 중첩 '{', 톱레벨 '}' 종료(i,true), 첫 '{' 전 '}' 무시(false), '"' 문자열 진입(inString 위임)까지 커버.

package ccnewsquest

import "testing"

func TestJSONScanStateFeed(t *testing.T) {
	// '}' before any '{' is ignored (start stays -1, no close).
	t.Run("close before open ignored", func(t *testing.T) {
		st := jsonScanState{start: -1}
		if _, done := st.feed('}', 0); done {
			t.Fatal("'}' before any '{' should not close")
		}
		if st.start != -1 {
			t.Fatalf("start = %d, want -1", st.start)
		}
	})

	// A full balanced pass: { { } } closes the top-level object at the last '}'.
	t.Run("balanced nested closes at top", func(t *testing.T) {
		bytes := []byte("{{}}")
		st := jsonScanState{start: -1}
		var end int
		var done bool
		for i, c := range bytes {
			end, done = st.feed(c, i)
			if done {
				break
			}
		}
		if !done {
			t.Fatal("expected the top-level object to close")
		}
		if st.start != 0 || end != 3 {
			t.Fatalf("start=%d end=%d, want 0 and 3", st.start, end)
		}
	})

	// A '"' flips into string mode → the next '{' is delegated to feedInString and
	// does not change depth.
	t.Run("quote enters string mode", func(t *testing.T) {
		st := jsonScanState{start: -1}
		st.feed('{', 0)
		st.feed('"', 1)
		if !st.inString {
			t.Fatal("expected inString after unescaped quote")
		}
		depthBefore := st.depth
		if _, done := st.feed('{', 2); done {
			t.Fatal("'{' inside string should not close anything")
		}
		if st.depth != depthBefore {
			t.Fatalf("depth changed inside string: %d → %d", depthBefore, st.depth)
		}
	})
}
