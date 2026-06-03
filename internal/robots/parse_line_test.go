//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what parseLine이 정상/주석제거/주석줄/빈줄/콜론없음/빈필드/소문자화를 올바르게 분리하는지 테이블로 검증한다.

package robots

import "testing"

func TestParseLine(t *testing.T) {
	tests := []struct {
		name      string
		line      string
		wantField string
		wantValue string
		wantOK    bool
	}{
		{"normal", "User-agent: *", "user-agent", "*", true},
		{"comment stripped", "Disallow: /x # note", "disallow", "/x", true},
		{"whole line comment", "# hello", "", "", false},
		{"blank", "   ", "", "", false},
		{"no colon", "garbage", "", "", false},
		{"empty field before colon", ": value", "", "", false},
		{"lowercases field", "DISALLOW: /Y", "disallow", "/Y", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, v, ok := parseLine(tt.line)
			if f != tt.wantField || v != tt.wantValue || ok != tt.wantOK {
				t.Errorf("parseLine(%q) = (%q,%q,%v), want (%q,%q,%v)",
					tt.line, f, v, ok, tt.wantField, tt.wantValue, tt.wantOK)
			}
		})
	}
}
