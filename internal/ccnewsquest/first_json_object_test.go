//ff:func feature=event6 type=helper control=iteration dimension=1
//ff:what firstJSONObject 경계·견고성 표 테스트. 순수객체/산문혼합/문자열내중괄호/이스케이프따옴표·백슬래시/객체없음/선행닫힘괄호·불균형까지 모든 분기 커버.

package ccnewsquest

import "testing"

func TestFirstJSONObject(t *testing.T) {
	cases := []struct {
		name   string
		in     string
		want   string
		wantOk bool
	}{
		{"pure", `{"a":1}`, `{"a":1}`, true},
		{"prose around", `prefix {"a":1} suffix`, `{"a":1}`, true},
		{"nested", `x {"a":{"b":2}} y`, `{"a":{"b":2}}`, true},
		// Braces inside a string literal must not affect depth counting.
		{"braces in string", `{"a":"x}y{z"}`, `{"a":"x}y{z"}`, true},
		// An escaped quote keeps the scanner inside the string; the trailing brace in
		// the value (\\) and \" exercise the escape branch.
		{"escaped quote and backslash", `{"a":"q\"} \\ {"}`, `{"a":"q\"} \\ {"}`, true},
		{"no object", `just prose, no braces`, "", false},
		// A '}' appearing before any '{' is skipped (start < 0 branch).
		{"close before open", `} then {"a":1}`, `{"a":1}`, true},
		// Braces never balance: open with no matching close.
		{"unbalanced open", `{"a":1`, "", false},
	}
	for _, c := range cases {
		got, ok := firstJSONObject(c.in)
		if ok != c.wantOk || got != c.want {
			t.Errorf("%s: firstJSONObject(%q) = (%q, %v), want (%q, %v)",
				c.name, c.in, got, ok, c.want, c.wantOk)
		}
	}
}
