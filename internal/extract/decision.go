//ff:type feature=extract type=model
//ff:what 신뢰 게이트 판정 결과와 임계 상수. PASS/SKIPPED 결정, SKIP 사유 문자열, 본문 텍스트 최소 길이 N(MinBodyLen).

package extract

// MinBodyLen is the minimum anchor-target body-text length (in bytes) for the
// trust gate to PASS. Below this the article is SKIPPED as "body too short for
// anchoring".
//
// Rationale: Phase006 anchors AI-asserted who/when/what tokens as substrings of
// this text. A few-sentence stub (nav blurbs, paywall teasers, link lists) gives
// no reliable substring surface and produces meaningless anchors. ~200 bytes is
// roughly two to three sentences of running prose — the floor below which there
// is not enough text to anchor six elements against. It is intentionally low
// (a permissive floor, not a quality bar): the goal is to reject empty/stub
// pages, not to judge article quality. For multibyte scripts (Korean, etc.) 200
// UTF-8 bytes is well under 200 characters, so it does not over-reject.
const MinBodyLen = 200

// SKIP reason strings (Phase005 design §trust gate). Recorded in
// Article.SkipReason when the gate does not PASS.
const (
	SkipNoStructured = "구조화 데이터 없음(JSON-LD/OG) — 신뢰 불가"
	SkipBodyTooShort = "본문 텍스트 부족 — 앵커 불가"
)

// Decision is the deterministic trust-gate outcome for one article.
type Decision struct {
	Pass       bool
	SkipReason string // "" when Pass; one of the Skip* constants otherwise
	Source     string // "jsonld" | "og" | "" — structured source that fed the gate
	BodyLen    int    // length (bytes) of the anchor-target body text
}
