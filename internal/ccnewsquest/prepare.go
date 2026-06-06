//ff:func feature=gate type=helper control=sequence level=error
//ff:what Definition.Prepare(s, it, raw). raw JSON을 session.Event6로 디코드하고 it.Payload(*session.Article)를 푼다. ① robots(A): host robots를 pick-time에 host당 1회 평가(s.Meta robots 캐시 적재, submit Save로 영속)해 거부면 OutBlock short verdict로 BLOCKED 단락(seed-time block_article의 OutBlock을 pick-time으로 이동). ② 본문(B/C): cacheDir·UA를 s.Meta에서 소싱(미기록 시 리시버 fallback)해 공유 헬퍼 readArticleBody로 WARC 재독+extract.Apply. 신뢰 게이트 탈락이면 SkipReason 되쓰고 OutSkip short verdict. ③ anchored(D): 통과면 fillAnchored로 각 present 필드 anchored를 게이트 동일 함수로 채운 뒤 a.Event6=&ev, SetPayload로 Extracted+Event6 되쓰고 Context{Item, Submission:&ev, Source:bodyText} 반환. WARC 디스크 IO는 결정론적(로컬 캐시).
package ccnewsquest

import (
	"encoding/json"
	"fmt"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/gate"
	"github.com/park-jun-woo/reins/pkg/quest"
)

// Prepare decodes the raw event6 JSON, evaluates robots at pick time, re-reads the
// article body from its WARC locator (a deterministic local-cache disk read, not a
// live fetch), runs the trust gate, stamps the per-field anchored verdict, and writes
// the enriched snapshot back into the item payload so the verified facts survive into
// the result JSONL (Phase012). It receives the session so it can source cacheDir/UA
// from Meta (Phase013 B) and read/update the per-host robots cache in Meta (Phase013
// A); submit Saves after Prepare, so those Meta writes persist.
//
//   - robots deny → an OutBlock short verdict locks the item to BLOCKED, moving the
//     former seed-time block (block_article) to pick time (Phase013 A).
//   - trust FAIL (no structured data / body too short) → the Article carries the
//     SkipReason; SetPayload keeps it, then an OutSkip verdict short-circuits the
//     anchor gate and the item locks to SKIPPED.
//   - trust PASS → the per-field anchored flags are filled with the gate's own
//     checkField (Phase013 D), the submitted event6 is attached, SetPayload writes
//     Extracted+Event6 back, and Prepare returns the anchor-evaluation Context.
func (d ccnewsDef) Prepare(s *quest.Session, it *quest.Item, raw []byte) (gate.Context, *quest.Verdict, error) {
	var ev session.Event6
	if err := json.Unmarshal(raw, &ev); err != nil {
		return gate.Context{}, nil, err
	}

	var a session.Article
	if err := it.DecodePayload(&a); err != nil {
		return gate.Context{}, nil, fmt.Errorf("prepare: item payload decode: %w", err)
	}

	userAgent, cacheDir := d.sourceConfig(s)

	// A. Pick-time robots: evaluate the host once and cache the result in Meta. A
	// denied host short-circuits to BLOCKED here (same OutBlock the seed used to set).
	if ok, reason := d.robots.allowed(s, userAgent, &a); !ok {
		a.State = session.BLOCKED
		a.SkipReason = reason
		_ = it.SetPayload(&a)
		return gate.Context{}, &quest.Verdict{
			Outcome: quest.OutBlock,
			Facts:   []quest.Fact{{Rule: "robots", Where: a.Host, Actual: reason}},
		}, nil
	}

	bodyText, ok, err := readArticleBody(userAgent, cacheDir, &a)
	if err != nil {
		return gate.Context{}, nil, err
	}
	if !ok {
		// Trust gate failed → SKIPPED short-circuit: the anchor gate never runs.
		return gate.Context{}, skipVerdict(it, &a), nil
	}

	// D. Stamp each present field's anchored flag with the gate's own checkField over
	// the same Source before the payload is written, so the artifact's anchored equals
	// the verdict field-for-field.
	fillAnchored(&ev, bodyText)

	// Trust PASS: attach the verified event6 and write the enriched snapshot
	// (Extracted + Event6) back into the payload before the anchor gate runs.
	a.Event6 = &ev
	if err := it.SetPayload(&a); err != nil {
		return gate.Context{}, nil, fmt.Errorf("prepare: payload 되쓰기: %w", err)
	}
	return gate.Context{Item: it, Submission: &ev, Source: bodyText}, nil, nil
}
