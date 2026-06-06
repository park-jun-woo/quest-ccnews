//ff:func feature=cli type=command control=sequence level=error
//ff:what agent 통합 (2) 펜스 두른 JSON→구제 PASS. 스텁이 ```json 펜스로 감싼 충실한 event6를 반환해도 Prepare 관용 디코드(Phase015 A)로 구제되어 1회 호출로 PASS(A 회귀 가드).

package main

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/reins/pkg/llm"
)

func TestAgentFencedRescue(t *testing.T) {
	sessionPath, outPath, cacheDir := writeAgentFixture(t)

	calls := 0
	backend := llm.CallFunc(func(system, user string) (string, error) {
		calls++
		return fencedGoodEvent6, nil
	})

	out := runAgent(t, sessionPath, outPath, cacheDir, backend)
	if calls != 1 {
		t.Fatalf("backend called %d times, want 1 (fenced JSON rescued, no retry)", calls)
	}
	if !strings.Contains(out, "PASS") {
		t.Fatalf("agent output = %q, want a PASS (fence rescued by lenient decode)", out)
	}
}
