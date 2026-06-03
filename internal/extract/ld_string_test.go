//ff:func feature=extract type=helper control=iteration dimension=1
//ff:what ldString이 문자열/배열첫원소/중첩배열/빈배열/숫자/nil/맵을 올바르게 강제변환하는지 테이블로 검증한다.

package extract

import "testing"

func TestLDString(t *testing.T) {
	cases := []struct {
		name string
		in   any
		want string
	}{
		{"string", "  hello  ", "hello"},
		{"array first string", []any{" first ", "second"}, "first"},
		{"array first nested array", []any{[]any{"deep"}}, "deep"},
		{"empty array", []any{}, ""},
		{"number", 42.0, ""},
		{"nil", nil, ""},
		{"map", map[string]any{"name": "x"}, ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := ldString(c.in); got != c.want {
				t.Fatalf("ldString(%v) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}
