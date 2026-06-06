//ff:func feature=event6 type=helper control=selection
//ff:what (*jsonScanState).feed(c, i) — 문자열 밖에서 한 바이트를 소비해 균형괄호 상태를 갱신하고, 이 바이트가 톱레벨 객체를 닫으면 (i, true)를 반환. 문자열 내부면 feedInString에 위임. '{'에서 첫 위치 기록+depth 증가, '}'에서 depth 감소(첫 '{' 전 닫힘은 무시), depth 0 도달 시 종료. firstJSONObject 루프 중첩을 ≤2로 낮추기 위한 추출 — 동작은 추출 전과 동일.
package ccnewsquest

// feed consumes one byte c at index i. While inside a string it delegates to
// feedInString. Otherwise it tracks brace state: '{' records the first start position
// and increments depth, '}' decrements depth (a '}' before any '{' is skipped), and
// reaching depth 0 closes the top-level object. It returns (i, true) on close and
// (0, false) otherwise, matching firstJSONObject's pre-extraction switch exactly.
func (st *jsonScanState) feed(c byte, i int) (int, bool) {
	if st.inString {
		st.feedInString(c)
		return 0, false
	}
	switch c {
	case '"':
		st.inString = true
	case '{':
		if st.start < 0 {
			st.start = i
		}
		st.depth++
	case '}':
		if st.start < 0 {
			return 0, false
		}
		st.depth--
		if st.depth == 0 {
			return i, true
		}
	}
	return 0, false
}
