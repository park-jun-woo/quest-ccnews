//ff:type feature=output type=model
//ff:what 출력 event6의 한 요소 — value(영어)와 anchors(원어, 없으면 생략).

package output

// OutField is one event6 element in the output: value (English) + anchors
// (original language, omitted when none).
type OutField struct {
	Value   string   `json:"value"`
	Anchors []string `json:"anchors,omitempty"`
}
