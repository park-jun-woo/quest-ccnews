//ff:type feature=event6 type=model
//ff:what 육하원칙 6요소 추출(누가·언제·어디서·무엇·어떻게·왜). 값은 영어, 앵커는 원문 토큰. 없으면 nil.

package session

// Event6: the six-element (육하원칙) extraction. value is English, anchors are
// original-language tokens (Phase006). Absent fields are nil.
type Event6 struct {
	Who   *Field `json:"who,omitempty"`
	When  *Field `json:"when,omitempty"`
	Where *Field `json:"where,omitempty"`
	What  *Field `json:"what,omitempty"`
	How   *Field `json:"how,omitempty"`
	Why   *Field `json:"why,omitempty"`
}
