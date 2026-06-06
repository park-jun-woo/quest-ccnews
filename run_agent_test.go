//ff:func feature=cli type=helper control=sequence level=error
//ff:what runAgent(t, sessionPath, outPath, cacheDir, backend) — 주입한 llm.Backend 스텁으로 reins agent 루프(cli.NewQuestCmd+ccnewsquest.Def)를 픽스처 위에서 구동해 합쳐진 명령 출력을 돌려준다. Execute가 Go 에러를 반환하면 테스트 실패.
package main

import (
	"bytes"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/ccnewsquest"
	"github.com/park-jun-woo/reins/pkg/cli"
	"github.com/park-jun-woo/reins/pkg/llm"
)

// runAgent drives the reins agent loop over the fixture with an injected stub LLM
// backend, returning the combined command output.
func runAgent(t *testing.T, sessionPath, outPath, cacheDir string, backend llm.Backend) string {
	t.Helper()
	def := ccnewsquest.Def(defaultUserAgent, cacheDir)
	opts := cli.Options{Agent: &cli.AgentOptions{
		System:     ccnewsSystem,
		RuleSystem: ccnewsRuleCoaching,
		LLM:        backend,
	}}
	cmd := cli.NewQuestCmd("ccnews", def, opts)
	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&out)
	cmd.SetArgs([]string{"--session", sessionPath, "--out", outPath, "agent"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("agent execute: %v\noutput:\n%s", err, out.String())
	}
	return out.String()
}
