//ff:func feature=gate type=helper control=sequence level=error
//ff:what 신뢰 게이트 탈락 시 호출. 방금 기록된 SkipReason을 it.SetPayload로 산출물에 되쓴 뒤(best-effort), OutSkip short verdict(structured-trust Fact)를 만들어 돌려준다. 앵커 게이트를 단락한다.

package ccnewsquest

import (
	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

// skipVerdict handles the trust-gate FAIL path: the anchor gate never runs. It
// persists the just-recorded SkipReason so the SKIP cause survives into the
// output payload (best-effort; the OutSkip verdict is the source of truth), then
// returns an OutSkip short verdict carrying the structured-trust Fact.
func skipVerdict(it *quest.Item, a *session.Article) *quest.Verdict {
	_ = it.SetPayload(a)
	return &quest.Verdict{
		Outcome: quest.OutSkip,
		Facts: []quest.Fact{{
			Rule:   "structured-trust",
			Where:  "body",
			Actual: a.SkipReason,
		}},
	}
}
