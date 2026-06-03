//ff:func feature=output type=helper control=iteration dimension=1
//ff:what Render의 라우팅(nil 기사→nil, 비종단 상태 TODO/미지→nil)을 검증한다. 종단 분기(PASS/REVIEW/BLOCKED/SKIPPED/DONE)는 render_collected/render_audit 테스트에서 다룬다.

package output

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestRender(t *testing.T) {
	if got := Render(nil, nil); got != nil {
		t.Errorf("Render(nil) = %+v, want nil", got)
	}
	for _, st := range []session.State{session.TODO, session.State("WEIRD")} {
		if got := Render(&session.Article{State: st}, nil); got != nil {
			t.Errorf("Render(state=%s) = %+v, want nil", st, got)
		}
	}
}
