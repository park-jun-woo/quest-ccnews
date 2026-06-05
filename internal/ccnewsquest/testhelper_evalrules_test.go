//ff:func feature=gate type=helper control=sequence
//ff:what 규칙·동등성 테스트 공용 헬퍼 evalRules(ev,source)→gate.Evaluate(Rules) Outcome. submit가 타는 verdict 경로를 그대로 돌려준다.

package ccnewsquest

import (
	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"
)

// evalRules runs the full Rules() catalog over (ev, source) via gate.Evaluate and
// returns the aggregated Outcome — the verdict path ccnews's submit takes.
func evalRules(ev *session.Event6, source string) quest.Outcome {
	return gate.Evaluate(ccnewsDef{}.Rules(), ctxOf(ev, source)).Outcome
}
