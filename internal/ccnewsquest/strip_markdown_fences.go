//ff:func feature=event6 type=helper control=sequence
//ff:what stripMarkdownFences — LLM 출력에서 감싸진 markdown 코드 펜스를 벗긴다(yongol 동일 구현 이식).

package ccnewsquest

import "strings"

// stripMarkdownFences removes wrapping markdown code fences from LLM output. It is a
// verbatim port of yongol's strip_markdown_fences so a small local model's ```json …
// ``` envelope is shed before the event6 JSON is scanned. Ported, same behavior.
func stripMarkdownFences(s string) string {
	s = strings.TrimSpace(s)
	if !strings.HasPrefix(s, "```") {
		return s
	}
	// Remove opening fence (``` or ```json etc.)
	idx := strings.Index(s, "\n")
	if idx < 0 {
		return s
	}
	s = s[idx+1:]

	// Remove closing fence
	if last := strings.LastIndex(s, "```"); last >= 0 {
		s = s[:last]
	}
	return strings.TrimSpace(s)
}
