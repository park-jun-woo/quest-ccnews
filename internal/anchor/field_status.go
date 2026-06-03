//ff:type feature=anchor type=model
//ff:what 한 필드의 앵커 대조 결과 분류. anchored(모든 유효앵커가 원문 substring), unanchored(유효앵커 0개·구조적 미검증), hallucination(유효앵커 하나라도 원문에 없음). 유효앵커는 normalize 후 비지 않은 앵커.

package anchor

// fieldStatus is the per-field outcome of anchor checking. A "valid" anchor is
// one whose normalized form is non-empty (Phase009 L0); empty/whitespace anchors
// are ignored for both counting and matching.
type fieldStatus int

const (
	anchored      fieldStatus = iota // every valid anchor is a source substring (≥1 valid anchor)
	unanchored                       // zero valid (non-empty/non-whitespace) anchors — structurally unverifiable (REVIEW for optional)
	hallucination                    // some valid anchor is absent from the source (FAIL)
)
