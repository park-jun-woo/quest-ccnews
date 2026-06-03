//ff:func feature=anchor type=helper control=selection
//ff:what 필수 필드(who/when/what) 하나를 검사한다. nil이면 FAIL(값 없음), value가 무효(빈·짧음·플레이스홀더)면 FAIL, 앵커가 환각이거나 유효앵커 0개면 FAIL. 통과하면 nil. 게이트의 필수 루프 본문을 동작보존 추출한 것(순수, Field.Anchored 채움).

package anchor

import "fmt"

// checkRequired evaluates one required field. A nil field is FAIL ("부재" 변명
// 불가); an intrinsically invalid value (empty, too short, or a placeholder
// token — Phase009 L3) is FAIL; missing valid anchors or a hallucinated anchor is
// FAIL. It returns a FAIL *Result on failure, or nil when the field passes. Pure;
// checkField fills f.Anchored. Extracted from Gate's required loop to keep nesting
// at depth 2.
func checkRequired(nf namedField, normSource string) *Result {
	if nf.f == nil {
		return &Result{Verdict: FAIL, Reason: fmt.Sprintf("필수 필드 %s 누락(값 없음)", nf.name)}
	}
	if !validValue(nf.f.Value) {
		return &Result{Verdict: FAIL, Reason: fmt.Sprintf("필수 필드 %s 값이 플레이스홀더/공허함: %q", nf.name, nf.f.Value)}
	}
	switch st, bad := checkField(nf.f, normSource); st {
	case hallucination:
		return &Result{Verdict: FAIL, Reason: fmt.Sprintf("필수 필드 %s 앵커가 원문에 없음(환각): %q", nf.name, bad)}
	case unanchored:
		return &Result{Verdict: FAIL, Reason: fmt.Sprintf("필수 필드 %s 앵커 없음(검증 불가)", nf.name)}
	}
	return nil
}
