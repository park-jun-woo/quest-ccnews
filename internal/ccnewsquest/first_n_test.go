//ff:func feature=event6 type=helper control=iteration dimension=1
//ff:what firstN 경계 테스트. 입력이 n보다 짧음(전체 반환)/같음/김(앞 n바이트만)/빈 입력 모든 분기 커버.

package ccnewsquest

import "testing"

func TestFirstN(t *testing.T) {
	cases := []struct {
		name string
		in   string
		n    int
		want string
	}{
		{"shorter than n", "abc", 5, "abc"},
		{"equal to n", "abc", 3, "abc"},
		{"longer than n", "abcdef", 3, "abc"},
		{"empty input", "", 4, ""},
		{"zero n", "abc", 0, ""},
	}
	for _, c := range cases {
		if got := firstN([]byte(c.in), c.n); got != c.want {
			t.Errorf("%s: firstN(%q, %d) = %q, want %q", c.name, c.in, c.n, got, c.want)
		}
	}
}
