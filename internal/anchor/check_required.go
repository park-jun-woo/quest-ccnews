//ff:func feature=anchor type=helper control=selection
//ff:what 필수 필드(who/when/what) 하나를 검사한다. nil이거나 Value가 비면 FAIL, 앵커가 환각이거나 0개면 FAIL. 통과하면 nil. 게이트의 필수 루프 본문을 동작보존 추출한 것(순수, Field.Anchored 채움).

package anchor

import "fmt"

// checkRequired evaluates one required field. A missing value or missing anchors
// is FAIL ("부재" 변명 불가); a hallucinated anchor is FAIL. It returns a FAIL
// *Result on failure, or nil when the field passes. Pure; checkField fills
// f.Anchored. Extracted from Gate's required loop to keep nesting at depth 2.
func checkRequired(nf namedField, normSource string) *Result {
	if nf.f == nil || nf.f.Value == "" {
		return &Result{Verdict: FAIL, Reason: fmt.Sprintf("필수 필드 %s 누락(값 없음)", nf.name)}
	}
	switch st, bad := checkField(nf.f, normSource); st {
	case hallucination:
		return &Result{Verdict: FAIL, Reason: fmt.Sprintf("필수 필드 %s 앵커가 원문에 없음(환각): %q", nf.name, bad)}
	case unanchored:
		return &Result{Verdict: FAIL, Reason: fmt.Sprintf("필수 필드 %s 앵커 없음(검증 불가)", nf.name)}
	}
	return nil
}
