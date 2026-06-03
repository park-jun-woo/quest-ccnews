//ff:type feature=anchor type=model
//ff:what 한 필드의 앵커 대조 결과 분류. anchored(모든 앵커가 원문 substring), unanchored(앵커 0개·구조적 미검증), hallucination(앵커 하나라도 원문에 없음).

package anchor

// fieldStatus is the per-field outcome of anchor checking.
type fieldStatus int

const (
	anchored      fieldStatus = iota // every anchor is a source substring (Anchors non-empty)
	unanchored                       // len(Anchors)==0 — structurally unverifiable (REVIEW for optional)
	hallucination                    // some anchor is absent from the source (FAIL)
)
