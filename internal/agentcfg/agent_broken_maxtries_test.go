//ff:func feature=cli type=command control=sequence level=error
//ff:what agent 통합 (4) 계속 깨진 출력→MaxTries DONE(루프 중단 아님). 스텁이 매번 JSON 아닌 산문을 반환 → Prepare가 retryable FAIL short verdict(event6-json)로 흡수 → 래칫이 MaxTries 소진 후 DONE 잠금, 루프는 Go 에러 없이 정상 종료(Phase015 A 핵심: 한 번의 포맷 잡음이 무인 루프를 죽이지 않음).

package agentcfg

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/reins/pkg/llm"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestAgentBrokenMaxTries(t *testing.T) {
	sessionPath, outPath, cacheDir := writeAgentFixture(t)

	calls := 0
	backend := llm.CallFunc(func(system, user string) (string, error) {
		calls++
		return "I'm sorry, I can't produce that.", nil // never valid JSON
	})

	// runAgent fails the test if Execute returns a Go error; reaching here proves
	// the loop did NOT abort on format noise.
	out := runAgent(t, sessionPath, outPath, cacheDir, backend)
	if calls != quest.MaxTries {
		t.Fatalf("backend called %d times, want MaxTries=%d (then DONE locks the item)", calls, quest.MaxTries)
	}
	if !strings.Contains(out, "FAIL") {
		t.Fatalf("agent output = %q, want repeated FAIL", out)
	}

	// The item must end terminal (DONE) so the loop drains rather than spinning.
	s, err := quest.Load(sessionPath)
	if err != nil {
		t.Fatalf("Load session: %v", err)
	}
	if len(s.Items) != 1 || s.Items[0].State != quest.DONE {
		t.Fatalf("item state = %v, want DONE after MaxTries", s.Items[0].State)
	}
}
