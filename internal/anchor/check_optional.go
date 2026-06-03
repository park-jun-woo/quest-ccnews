//ff:func feature=anchor type=helper control=selection
//ff:what 선택 필드(where/how/why) 하나를 검사한다. nil이면 무시, value가 무효(빈·짧음·플레이스홀더)면 FAIL, 앵커 환각이면 FAIL, 유효앵커 0개면 미검증(REVIEW 후보). 게이트의 선택 루프 본문을 동작보존 추출한 것(순수, Field.Anchored 채움).

package anchor

import "fmt"

// checkOptional evaluates one optional field. A nil field is ignored. A present
// field with an intrinsically invalid value (empty, too short, or a placeholder
// token — Phase009 L3) is FAIL, checked before the anchor test. A hallucinated
// anchor is FAIL (returns a *Result). A present field with no valid anchors is
// structurally unverifiable (returns nil, true) — a REVIEW candidate. Anchored or
// absent fields return (nil, false). Pure; checkField fills f.Anchored. Extracted
// from Gate's optional loop to keep nesting at depth 2.
func checkOptional(nf namedField, normSource string) (*Result, bool) {
	if nf.f == nil {
		return nil, false
	}
	if !validValue(nf.f.Value) {
		return &Result{Verdict: FAIL, Reason: fmt.Sprintf("선택 필드 %s 값이 플레이스홀더/공허함: %q", nf.name, nf.f.Value)}, false
	}
	switch st, bad := checkField(nf.f, normSource); st {
	case hallucination:
		return &Result{Verdict: FAIL, Reason: fmt.Sprintf("선택 필드 %s 앵커가 원문에 없음(환각): %q", nf.name, bad)}, false
	case unanchored:
		return nil, true
	}
	return nil, false
}
