//ff:func feature=gate type=helper control=sequence
//ff:what 규칙·동등성 테스트 공용 헬퍼 ctxOf(ev,source)→앵커평가용 gate.Context.

package ccnewsquest

import (
	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/gate"
)

// ctxOf builds the anchor-evaluation Context for an event6 against a source.
func ctxOf(ev *session.Event6, source string) gate.Context {
	return gate.Context{Submission: ev, Source: source}
}
