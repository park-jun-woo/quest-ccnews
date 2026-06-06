//ff:func feature=cli type=config control=sequence level=error
//ff:what ccnews 에이전트 루프의 고정 설정값. robots UA(UserAgent), WARC 캐시 디렉터리(CacheDir), 시스템 프롬프트(System), 규칙별 코칭 되먹임(RuleCoaching)을 main과 agent 통합 테스트가 공유한다 — "생성은 AI, 판정은 결정론적 앵커 게이트".

package agentcfg

// UserAgent matches Phase001 결정 4 (robots UA).
const UserAgent = "parkjunwoo-quest/0.1 (+https://www.parkjunwoo.com)"

// CacheDir is where downloaded .warc.gz files are cached on disk; it is the
// same directory Prepare's WARC re-read client reads from, keeping run and submit
// consistent (Phase013).
const CacheDir = "warc-cache"

// System is the global system prompt for the agent loop: it states that the
// model is the generator and a deterministic anchor gate is the judge — anchors must
// be verbatim original-text tokens, values are English (and not judged).
const System = "You extract the six W's (who/what/when/where/how/why) from a news " +
	"article and output ONLY a single event6 JSON object — no prose, no markdown. " +
	"Each field has a 'value' (English; dates ISO, numbers normalized) and 'anchors' " +
	"(an array of tokens that appear VERBATIM in the original article text). " +
	"who and what are required; when/where/how/why are optional (null when absent). " +
	"A deterministic gate checks only that every anchor is an exact substring of the " +
	"original body. Never invent an anchor that is not literally in the text."

// RuleCoaching maps each anchor-gate rule ID (ccnewsquest rules.go) to extra
// system guidance applied on the next retry when the previous attempt FAILed on it.
var RuleCoaching = map[string]string{
	"event6-json": "Output exactly one JSON object and nothing else — no fences, no commentary.",
	"required-present": "The previous attempt left a required field (who/what) empty. " +
		"Provide a non-empty value AND at least one anchor for both who and what.",
	"required-anchor-valid": "A required field had a value but no anchors, or malformed anchors. " +
		"Give each required field at least one anchor that is a verbatim substring of the article.",
	"required-anchor-real": "A required field's anchor was NOT found in the article (hallucination). " +
		"Use only tokens copied exactly from the original body as anchors.",
	"optional-present": "An optional field had anchors but no value. Either give it a value or set it to null.",
	"optional-anchor-real": "An optional field's anchor was NOT found in the article (hallucination). " +
		"Drop that field to null, or use only verbatim tokens as anchors.",
}
