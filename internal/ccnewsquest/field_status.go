//ff:type feature=gate type=model
//ff:what 한 필드의 앵커 대조 결과 분류(anchor.fieldStatus의 이식). statusAnchored(유효앵커 전부 원문 substring, ≥1), statusUnanchored(유효앵커 0개·구조적 미검증), statusHallucination(유효앵커 하나라도 원문에 없음). 유효앵커는 textmatch.Normalize 후 비지 않은 앵커.

package ccnewsquest

// fieldStatus is the per-field outcome of anchor checking — the reins-port of
// anchor.fieldStatus. A "valid" anchor is one whose textmatch.Normalize form is
// non-empty (Phase009 L0); empty/whitespace anchors are ignored for both counting
// and matching.
type fieldStatus int

const (
	statusAnchored      fieldStatus = iota // every valid anchor is a source substring (≥1 valid anchor)
	statusUnanchored                       // zero valid anchors — structurally unverifiable (REVIEW for optional)
	statusHallucination                    // some valid anchor is absent from the source (FAIL)
)
