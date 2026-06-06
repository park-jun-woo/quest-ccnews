//ff:func feature=cli type=command control=sequence level=error
//ff:what 프로그램 진입점. reins cli.NewQuestCmd로 ccnews 퀘스트 CLI(scan/next/submit/status/export/rules)를 조립하고, 투트랙 WARC 인제스천 `run` 명령을 ExtraCommand로 부착해 실행한다. 도메인 로직은 ccnewsquest.Def(userAgent, cacheDir) 하나, 인제스천은 runcmd.New() 하나.

package main

import (
	"os"

	"github.com/park-jun-woo/quest-ccnews/internal/ccnewsquest"
	"github.com/park-jun-woo/quest-ccnews/internal/runcmd"
	"github.com/park-jun-woo/reins/pkg/cli"
	"github.com/spf13/cobra"
)

// defaultUserAgent matches Phase001 결정 4 (robots UA).
const defaultUserAgent = "parkjunwoo-quest/0.1 (+https://www.parkjunwoo.com)"

// defaultCacheDir is where downloaded .warc.gz files are cached on disk; it is the
// same directory Prepare's WARC re-read client reads from, keeping run and submit
// consistent (Phase013).
const defaultCacheDir = "warc-cache"

// ccnewsSystem is the global system prompt for the agent loop: it states that the
// model is the generator and a deterministic anchor gate is the judge — anchors must
// be verbatim original-text tokens, values are English (and not judged).
const ccnewsSystem = "You extract the six W's (who/what/when/where/how/why) from a news " +
	"article and output ONLY a single event6 JSON object — no prose, no markdown. " +
	"Each field has a 'value' (English; dates ISO, numbers normalized) and 'anchors' " +
	"(an array of tokens that appear VERBATIM in the original article text). " +
	"who and what are required; when/where/how/why are optional (null when absent). " +
	"A deterministic gate checks only that every anchor is an exact substring of the " +
	"original body. Never invent an anchor that is not literally in the text."

// ccnewsRuleCoaching maps each anchor-gate rule ID (ccnewsquest rules.go) to extra
// system guidance applied on the next retry when the previous attempt FAILed on it.
var ccnewsRuleCoaching = map[string]string{
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

func main() {
	root := cli.NewQuestCmd("ccnews", ccnewsquest.Def(defaultUserAgent, defaultCacheDir), cli.Options{
		Version:       "0.3",
		ExtraCommands: []*cobra.Command{runcmd.New(defaultUserAgent, defaultCacheDir)},
		Agent: &cli.AgentOptions{
			System:     ccnewsSystem,
			RuleSystem: ccnewsRuleCoaching,
		},
	})
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
