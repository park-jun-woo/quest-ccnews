//ff:func feature=anchor type=helper control=iteration dimension=1
//ff:what normalize가 공백 런 접기·양끝 트림·유니코드 공백 처리를 케이스별로 검증한다.

package anchor

import "testing"

func TestNormalize(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", ""},
		{"only spaces", "   ", ""},
		{"trim ends", "  hello  ", "hello"},
		{"collapse runs", "a   b\t\tc", "a b c"},
		{"newlines and tabs", "a\n\tb\r\nc", "a b c"},
		{"unicode space", "a b　c", "a b c"},
		{"already normal", "a b c", "a b c"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := normalize(tc.in); got != tc.want {
				t.Errorf("normalize(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
