//ff:func feature=extract type=helper control=iteration dimension=1
//ff:what ldName이 문자열/객체(.name)/객체배열/문자열배열/빈배열/숫자/nil을 올바르게 이름으로 정규화하는지 테이블로 검증한다.

package extract

import "testing"

func TestLDName(t *testing.T) {
	cases := []struct {
		name string
		in   any
		want string
	}{
		{"string", " Jane Doe ", "Jane Doe"},
		{"object with name", map[string]any{"name": " Acme News "}, "Acme News"},
		{"object without name", map[string]any{"foo": "bar"}, ""},
		{"array first object", []any{map[string]any{"name": "A"}, map[string]any{"name": "B"}}, "A"},
		{"array first string", []any{"X"}, "X"},
		{"empty array", []any{}, ""},
		{"number", 7.0, ""},
		{"nil", nil, ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := ldName(c.in); got != c.want {
				t.Fatalf("ldName(%v) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}
