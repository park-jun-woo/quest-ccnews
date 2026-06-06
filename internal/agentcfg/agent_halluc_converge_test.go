//ff:func feature=cli type=command control=sequence level=error
//ff:what agent 통합 (3) 환각 앵커→FAIL 코칭 재시도→수렴. 스텁이 1회차에 원문에 없는 앵커(환각)를 반환 → required-anchor-real FAIL → 코칭 되먹임 → 2회차 충실 event6 → PASS 수렴.

package agentcfg

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/reins/pkg/llm"
)

func TestAgentHallucConverge(t *testing.T) {
	sessionPath, outPath, cacheDir := writeAgentFixture(t)

	calls := 0
	backend := llm.CallFunc(func(system, user string) (string, error) {
		calls++
		if calls == 1 {
			return hallucEvent6, nil // hallucinated anchor → FAIL
		}
		return goodEvent6, nil // retry is faithful → PASS
	})

	out := runAgent(t, sessionPath, outPath, cacheDir, backend)
	if calls != 2 {
		t.Fatalf("backend called %d times, want 2 (FAIL then PASS)", calls)
	}
	if !strings.Contains(out, "FAIL") || !strings.Contains(out, "PASS") {
		t.Fatalf("agent output = %q, want a FAIL then a PASS (convergence)", out)
	}
}
