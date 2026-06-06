//ff:func feature=event6 type=helper control=iteration dimension=1
//ff:what stripMarkdownFences 표 테스트. 펜스없음/```json 펜스/언어태그없는 ``` 펜스/닫는펜스없음/개행없는 펜스 모든 분기 커버.

package ccnewsquest

import "testing"

func TestStripMarkdownFences(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"no fence", `{"a":1}`, `{"a":1}`},
		{"json fence", "```json\n{\"a\":1}\n```", `{"a":1}`},
		{"bare fence", "```\n{\"a\":1}\n```", `{"a":1}`},
		// Opening fence present but no newline: returns input unchanged (idx < 0 branch).
		{"no newline after fence", "```json {\"a\":1}", "```json {\"a\":1}"},
		// Opening fence + newline but no closing fence: keep body after opening line.
		{"no closing fence", "```json\n{\"a\":1}", `{"a":1}`},
		// Leading/trailing whitespace is trimmed before fence detection.
		{"surrounding whitespace", "  ```json\n{\"a\":1}\n```  ", `{"a":1}`},
	}
	for _, c := range cases {
		if got := stripMarkdownFences(c.in); got != c.want {
			t.Errorf("%s: stripMarkdownFences(%q) = %q, want %q", c.name, c.in, got, c.want)
		}
	}
}
