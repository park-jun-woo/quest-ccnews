//ff:func feature=gate type=helper control=sequence
//ff:what Context.Submission을 *session.Event6로 풀고, anchor.Gate와 동일한 필수(who/what)·선택(when/where/how/why) namedField 리스트를 만든다. 비-Event6 제출이면 ok=false(규칙 no-fire). 규칙들이 공유하는 디코드 진입점.

package ccnewsquest

import (
	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/gate"
)

// event6Of unwraps the submission as a *session.Event6 and returns the required and
// optional field lists in the exact order anchor.Gate uses (who/what required;
// when/where/how/why optional). ok is false when the submission is not a non-nil
// *session.Event6, in which case a rule no-fires. This is the rules' shared decode
// entry point; the field ordering is what makes the per-rule "first offender"
// Facts match the original gate's first-return reason.
func event6Of(ctx gate.Context) (required, optional []namedField, ok bool) {
	ev, ok := ctx.Submission.(*session.Event6)
	if !ok || ev == nil {
		return nil, nil, false
	}
	required = []namedField{
		{"who", ev.Who}, {"what", ev.What},
	}
	optional = []namedField{
		{"when", ev.When}, {"where", ev.Where}, {"how", ev.How}, {"why", ev.Why},
	}
	return required, optional, true
}
