//ff:type feature=event6 type=model
//ff:what jsonScanState — firstJSONObject의 한-패스 균형괄호 스캐너 상태. start(첫 '{' 위치, 미발견 -1)·depth(중괄호 중첩)·inString(문자열 리터럴 내부)·escaped(직전 바이트가 백슬래시). 상태와 스캔 로직을 분리해 firstJSONObject 루프 중첩깊이를 ≤2로 유지하기 위한 추출 타입.
package ccnewsquest

// jsonScanState carries the single-pass brace-balancing state for firstJSONObject:
// the index of the first '{' (start, -1 until seen), the current brace nesting depth,
// whether the scanner is inside a JSON string literal, and whether the previous byte
// was an escaping backslash. It exists so the per-byte logic can live in a method
// (jsonScanState.feed) and keep firstJSONObject's loop nesting at depth ≤ 2.
type jsonScanState struct {
	start    int
	depth    int
	inString bool
	escaped  bool
}
