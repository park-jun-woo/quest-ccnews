//ff:func feature=gate type=helper control=sequence level=error
//ff:what Definition.Prepare. raw JSON을 session.Event6로 디코드하고, it.Payload(*session.Article)의 WARC를 ingest.Client.ReadBody로 재독→extract.Apply로 추출+신뢰 게이트. 신뢰 게이트 탈락(!ok)이면 OutSkip short verdict(structured-trust Fact, a.SkipReason)로 앵커 게이트 단락. 통과면 Context{Item, Submission:&ev, Source:bodyText}. WARC 디스크 IO는 결정론적(로컬 캐시, 라이브 fetch 아님).

package ccnewsquest

import (
	"encoding/json"
	"fmt"

	"github.com/park-jun-woo/quest-ccnews/internal/extract"
	"github.com/park-jun-woo/quest-ccnews/internal/ingest"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"
)

// Prepare decodes the raw event6 JSON and re-reads the article body from its WARC
// locator (ingest.Client.ReadBody — a deterministic local-cache disk read, not a
// live fetch), then runs extract.Apply (structured-field extraction + trust gate).
// When the trust gate fails (ok==false: no structured data or body too short),
// it short-circuits the anchor gate with an OutSkip verdict carrying a
// structured-trust Fact and the SkipReason extract.Apply just recorded — the item
// locks to SKIPPED without running the anchor rules. On trust PASS it returns the
// anchor-evaluation Context (Submission=&ev, Source=the anchor-target body text).
//
// userAgent/cacheDir come from the ccnewsDef receiver (Phase013 will source them
// from the session meta slot). This mirrors cmd/submit.go's WARC-reread flow.
func (d ccnewsDef) Prepare(it *quest.Item, raw []byte) (gate.Context, *quest.Verdict, error) {
	var ev session.Event6
	if err := json.Unmarshal(raw, &ev); err != nil {
		return gate.Context{}, nil, err
	}

	a, ok := it.Payload.(*session.Article)
	if !ok || a == nil {
		return gate.Context{}, nil, fmt.Errorf("prepare: item payload is not *session.Article")
	}

	client := ingest.NewClient(d.userAgent, d.cacheDir)
	htmlBytes, err := client.ReadBody(a.WARC)
	if err != nil {
		return gate.Context{}, nil, fmt.Errorf("원문 재독 실패 (%s): %w", a.URL, err)
	}

	bodyText, ok := extract.Apply(a, htmlBytes)
	if !ok {
		// Trust gate failed → SKIPPED short-circuit: the anchor gate never runs.
		return gate.Context{}, &quest.Verdict{
			Outcome: quest.OutSkip,
			Facts: []quest.Fact{{
				Rule:   "structured-trust",
				Where:  "body",
				Actual: a.SkipReason,
			}},
		}, nil
	}
	return gate.Context{Item: it, Submission: &ev, Source: bodyText}, nil, nil
}
