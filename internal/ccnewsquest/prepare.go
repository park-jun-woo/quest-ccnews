//ff:func feature=gate type=helper control=sequence level=error
//ff:what Definition.Prepare. raw JSON을 session.Event6로 디코드하고, it.Payload(*session.Article)의 WARC를 ingest.Client.ReadBody로 재독→extract.Apply로 추출+신뢰 게이트. 신뢰 게이트 탈락(!ok)이면 a.SkipReason을 it.SetPayload로 산출물에 되쓴 뒤 OutSkip short verdict(structured-trust Fact)로 앵커 게이트 단락. 통과면 a.Event6=&ev 부착 후 it.SetPayload로 Extracted+Event6를 payload에 되쓰고 Context{Item, Submission:&ev, Source:bodyText} 반환. enrich한 Article을 payload에 되씀(PASS=Event6+Extracted, SKIP=SkipReason). WARC 디스크 IO는 결정론적(로컬 캐시, 라이브 fetch 아님).

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
// After every mutation of the Article is done, Prepare writes the enriched
// snapshot back into the item payload with it.SetPayload so the verified facts
// actually survive into the result JSONL (Phase012). The reins submit flow
// (Prepare → Evaluate → quest.Apply → Save → Export) preserves the payload as-is,
// so Prepare is the only place the enrich can be persisted:
//
//   - trust FAIL (ok==false: no structured data or body too short) → the Article
//     already carries the SkipReason extract.Apply just recorded; SetPayload keeps
//     that reason in the output, then an OutSkip verdict (structured-trust Fact)
//     short-circuits the anchor gate and the item locks to SKIPPED.
//   - trust PASS → the submitted event6 is attached (a.Event6=&ev) and SetPayload
//     writes Extracted+Event6 back, then Prepare returns the anchor-evaluation
//     Context (Submission=&ev, Source=the anchor-target body text).
//
// userAgent/cacheDir come from the ccnewsDef receiver (Phase013 will source them
// from the session meta slot). This mirrors cmd/submit.go's WARC-reread flow.
func (d ccnewsDef) Prepare(it *quest.Item, raw []byte) (gate.Context, *quest.Verdict, error) {
	var ev session.Event6
	if err := json.Unmarshal(raw, &ev); err != nil {
		return gate.Context{}, nil, err
	}

	var a session.Article
	if err := it.DecodePayload(&a); err != nil {
		return gate.Context{}, nil, fmt.Errorf("prepare: item payload decode: %w", err)
	}

	client := ingest.NewClient(d.userAgent, d.cacheDir)
	htmlBytes, err := client.ReadBody(a.WARC)
	if err != nil {
		return gate.Context{}, nil, fmt.Errorf("원문 재독 실패 (%s): %w", a.URL, err)
	}

	bodyText, ok := extract.Apply(&a, htmlBytes)
	if !ok {
		// Trust gate failed → SKIPPED short-circuit: the anchor gate never runs.
		return gate.Context{}, skipVerdict(it, &a), nil
	}

	// Trust PASS: attach the verified event6 and write the enriched snapshot
	// (Extracted + Event6) back into the payload before the anchor gate runs.
	a.Event6 = &ev
	if err := it.SetPayload(&a); err != nil {
		return gate.Context{}, nil, fmt.Errorf("prepare: payload 되쓰기: %w", err)
	}
	return gate.Context{Item: it, Submission: &ev, Source: bodyText}, nil, nil
}
