//ff:func feature=anchor type=helper control=iteration dimension=1
//ff:what validValue가 빈·1룬·트림 후 1룬·플레이스홀더(대소문자 무시)를 무효로, 정상 토큰을 유효로 판정하는지 검증한다(Phase009 L3).

package anchor

import "testing"

func TestValidValue(t *testing.T) {
	invalid := []string{
		"", " ", "\t", "a", " a ", "-", "--",
		"Subject", "subject", "  Subject  ", "EVENT", "Event",
		"Unknown", "N/A", "n/a", "NA", "None", "null", "TBD",
	}
	for _, v := range invalid {
		if validValue(v) {
			t.Errorf("validValue(%q) = true, want false (junk)", v)
		}
	}

	valid := []string{
		"Alice", "UN", "은행", "署名", "12", "war", "Paris treaty",
	}
	for _, v := range valid {
		if !validValue(v) {
			t.Errorf("validValue(%q) = false, want true (real value)", v)
		}
	}
}
