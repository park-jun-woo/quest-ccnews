//ff:func feature=gate type=helper control=iteration dimension=1
//ff:what validValue 단위테스트. 트림 후 룬<2면 false, 플레이스홀더 블록리스트(대소문자 무시)면 false, 그 외 true. 빈/공백·단일룬·플레이스홀더·정상값·멀티바이트(룬길이) 분기 커버.

package ccnewsquest

import "testing"

func TestValidValue(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want bool
	}{
		{"empty", "", false},
		{"whitespace only", "   ", false},
		{"single rune", "a", false},
		{"single rune after trim", "  a  ", false},
		{"two runes ok", "ab", true},
		{"placeholder exact", "unknown", false},
		{"placeholder uppercase", "N/A", false},
		{"placeholder mixed case", "Subject", false},
		{"placeholder dash", "--", false},
		{"hyphen single is too short anyway", "-", false},
		{"normal value", "Alice", true},
		{"two-byte runes count as 2", "한국", true},
		{"trimmed normal", "  Paris  ", true},
		{"contains placeholder word but not exact", "unknown soldier", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := validValue(tc.in); got != tc.want {
				t.Fatalf("validValue(%q) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}
