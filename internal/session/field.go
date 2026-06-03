//ff:type feature=event6 type=model
//ff:what event6 한 요소. Anchored는 기계의 사실-앵커 판정. Interpretive는 표시용 라벨일 뿐 게이트 입력 아님.

package session

// Field: one event6 element. Anchored is the machine fact-anchor verdict.
type Field struct {
	Value    string   `json:"value"`
	Anchors  []string `json:"anchors,omitempty"`
	Anchored bool     `json:"anchored"`

	// Interpretive is a display label only — NOT a gate input. REVIEW is decided
	// by the machine from len(Anchors)==0, never from this flag (Phase006).
	Interpretive bool `json:"interpretive,omitempty"`
}
