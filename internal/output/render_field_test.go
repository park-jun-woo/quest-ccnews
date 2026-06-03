//ff:func feature=output type=helper control=sequence
//ff:what renderField가 nil은 nil로, 존재하는 Field는 value+anchors를 그대로 OutField로 옮기는지 검증한다.

package output

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestRenderField(t *testing.T) {
	if got := renderField(nil); got != nil {
		t.Errorf("renderField(nil) = %+v, want nil", got)
	}
	f := &session.Field{Value: "v", Anchors: []string{"a", "b"}}
	got := renderField(f)
	if got == nil || got.Value != "v" || !reflect.DeepEqual(got.Anchors, []string{"a", "b"}) {
		t.Errorf("renderField() = %+v", got)
	}
}
