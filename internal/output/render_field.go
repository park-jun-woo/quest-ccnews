//ff:func feature=output type=helper control=sequence
//ff:what session.Field 1건을 출력용 OutField(value+anchors)로 변환한다. nil이면 nil. 순수.

package output

import "github.com/park-jun-woo/quest-ccnews/internal/session"

// renderField converts one session Field into the output OutField (value +
// anchors), returning nil for an absent field.
func renderField(f *session.Field) *OutField {
	if f == nil {
		return nil
	}
	return &OutField{Value: f.Value, Anchors: f.Anchors}
}
