//ff:func feature=event6 type=helper control=iteration dimension=1
//ff:what firstJSONObject — 산문이 섞인 텍스트에서 첫 균형 JSON 객체({…매칭 }) 한 패스 깊이 카운트로 추출. 문자열 리터럴 내 중괄호·이스케이프 따옴표 무시. 바이트별 상태 전이는 jsonScanState.feed로 위임(중첩깊이 ≤2 유지).

package ccnewsquest

// firstJSONObject scans s for the first balanced top-level JSON object — from the
// first '{' to the '}' that closes it — and returns that slice. A single depth-counting
// pass (via jsonScanState.feed) tracks brace nesting while ignoring any '{'/'}' that
// appear inside a JSON string literal and following escape state so an escaped quote
// (\") inside a string does not falsely end the string. Returns ("", false) when no
// balanced object is present (no '{', or the braces never balance).
func firstJSONObject(s string) (string, bool) {
	st := jsonScanState{start: -1}
	for i := 0; i < len(s); i++ {
		end, done := st.feed(s[i], i)
		if done {
			return s[st.start : end+1], true
		}
	}
	return "", false
}
