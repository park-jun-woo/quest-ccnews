//ff:func feature=gate type=helper control=sequence level=error
//ff:what trimBody 단위 테스트. 본문 길이 상한(renderBodyMax) trim의 경계를 커버한다. ① 상한 이하·정확히 상한: 그대로 반환(생략 마커 없음). ② 상한 초과: 앞 renderBodyMax 바이트 + 생략 마커. ③ 빈 문자열: 그대로.
package ccnewsquest

import (
	"strings"
	"testing"
)

func TestTrimBody(t *testing.T) {
	const marker = "본문 일부 생략"

	t.Run("empty string is returned unchanged", func(t *testing.T) {
		if got := trimBody(""); got != "" {
			t.Fatalf("trimBody(\"\") = %q, want \"\"", got)
		}
	})

	t.Run("below cap is returned unchanged", func(t *testing.T) {
		in := strings.Repeat("a", renderBodyMax-1)
		got := trimBody(in)
		if got != in {
			t.Fatalf("short body altered (len %d → %d)", len(in), len(got))
		}
		if strings.Contains(got, marker) {
			t.Errorf("unexpected truncation marker on a short body")
		}
	})

	t.Run("exactly at cap is returned unchanged (boundary)", func(t *testing.T) {
		in := strings.Repeat("a", renderBodyMax)
		got := trimBody(in)
		if got != in {
			t.Fatalf("body of exactly renderBodyMax was altered")
		}
		if strings.Contains(got, marker) {
			t.Errorf("unexpected truncation marker at the exact cap")
		}
	})

	t.Run("above cap is truncated with a marker", func(t *testing.T) {
		in := strings.Repeat("a", renderBodyMax+100)
		got := trimBody(in)
		if !strings.HasPrefix(got, in[:renderBodyMax]) {
			t.Fatalf("truncated body does not start with the first renderBodyMax bytes")
		}
		if !strings.Contains(got, marker) {
			t.Fatalf("over-cap body missing truncation marker: %q", got[len(got)-60:])
		}
		// The kept text is exactly the first renderBodyMax bytes (the rest is marker).
		if len(got) <= renderBodyMax {
			t.Errorf("truncated output len=%d, want > renderBodyMax (kept bytes + marker)", len(got))
		}
	})
}
