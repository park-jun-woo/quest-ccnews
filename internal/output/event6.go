//ff:type feature=output type=model
//ff:what 출력용 event6 — 존재하는 필드만 영어 value와 원어 anchors(있을 때만)를 담는다. 부재 필드는 생략(omitempty).

package output

// Event6 is the output event6: each present field carries the English value and
// its original-language anchors (only when present). Absent fields are omitted.
type Event6 struct {
	Who   *OutField `json:"who,omitempty"`
	When  *OutField `json:"when,omitempty"`
	Where *OutField `json:"where,omitempty"`
	What  *OutField `json:"what,omitempty"`
	How   *OutField `json:"how,omitempty"`
	Why   *OutField `json:"why,omitempty"`
}
