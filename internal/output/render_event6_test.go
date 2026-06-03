//ff:func feature=output type=helper control=sequence
//ff:what renderEvent6가 nil은 nil로, 존재하는 필드만 매핑하고 부재 필드는 nil로 남기는지 검증한다.

package output

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestRenderEvent6(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		if got := renderEvent6(nil); got != nil {
			t.Errorf("renderEvent6(nil) = %+v, want nil", got)
		}
	})

	t.Run("present only", func(t *testing.T) {
		ev := &session.Event6{
			Who:  &session.Field{Value: "w"},
			When: &session.Field{Value: "t", Anchors: []string{"x"}},
		}
		out := renderEvent6(ev)
		if out == nil || out.Who == nil || out.When == nil {
			t.Fatalf("renderEvent6() = %+v", out)
		}
		if out.Where != nil || out.What != nil || out.How != nil || out.Why != nil {
			t.Errorf("absent fields should be nil: %+v", out)
		}
		if out.Who.Value != "w" {
			t.Errorf("Who.Value = %q", out.Who.Value)
		}
		if !reflect.DeepEqual(out.When.Anchors, []string{"x"}) {
			t.Errorf("When.Anchors = %v", out.When.Anchors)
		}
	})
}
