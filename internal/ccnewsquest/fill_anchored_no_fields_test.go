//ff:func feature=gate type=helper control=sequence
//ff:what fillAnchored 단위 테스트(부재 케이스). 전부 nil인 event6에 대해 fillAnchored가 패닉하지 않고 nil 필드를 그대로 둔다(present 필드만 채움).
package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestFillAnchoredNoFieldsPresent(t *testing.T) {
	// An all-nil event6 must not panic.
	ev := &session.Event6{}
	fillAnchored(ev, "any source")
	if ev.Who != nil || ev.What != nil {
		t.Fatalf("nil fields became non-nil: %+v", ev)
	}
}
