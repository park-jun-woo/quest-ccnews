//ff:func feature=event6 type=helper control=selection
//ff:what (*jsonScanState).feedInString(c) — 문자열 리터럴 내부에서 한 바이트를 소비해 이스케이프/종료 상태만 갱신. 직전이 백슬래시면 이 바이트는 이스케이프됨(\" \\), '\\'이면 다음을 이스케이프, '"'이면 문자열 종료. firstJSONObject 스캔 분리의 일부 — 동작은 추출 전 inString 분기와 동일.
package ccnewsquest

// feedInString consumes one byte c while the scanner is inside a JSON string literal,
// updating only escape/termination state: an escaped byte (after a backslash) is
// consumed as-is, a backslash arms the next escape, and an unescaped quote ends the
// string. Behavior is identical to firstJSONObject's pre-extraction inString branch.
func (st *jsonScanState) feedInString(c byte) {
	switch {
	case st.escaped:
		st.escaped = false
	case c == '\\':
		st.escaped = true
	case c == '"':
		st.inString = false
	}
}
