//ff:func feature=cli type=command control=sequence level=error
//ff:what agent 통합 (1) 정상→PASS. 스텁이 충실한 event6 JSON을 1회 반환 → 앵커 게이트 PASS, 1회 호출로 잠금, 결과 JSONL에 원문 URL+event6(사실) 기록(E 공개안전 불변 동반 단언).

package main

import (
	"os"
	"strings"
	"testing"

	"github.com/park-jun-woo/reins/pkg/llm"
)

func TestAgentPass(t *testing.T) {
	sessionPath, outPath, cacheDir := writeAgentFixture(t)

	calls := 0
	backend := llm.CallFunc(func(system, user string) (string, error) {
		calls++
		return goodEvent6, nil
	})

	out := runAgent(t, sessionPath, outPath, cacheDir, backend)
	if calls != 1 {
		t.Fatalf("backend called %d times, want 1 (clean PASS, no retry)", calls)
	}
	if !strings.Contains(out, "PASS") {
		t.Fatalf("agent output = %q, want a PASS", out)
	}

	jsonl, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatalf("read out.jsonl: %v", err)
	}
	rec := string(jsonl)
	// E (공개안전 불변): 결과는 원문 링크(URL)와 농축된 사실(event6)을 싣는다.
	if !strings.Contains(rec, "https://example.com/a") {
		t.Errorf("out.jsonl missing original URL link:\n%s", rec)
	}
	if !strings.Contains(rec, "event6") || !strings.Contains(rec, "Reporter") {
		t.Errorf("out.jsonl missing distilled event6 facts:\n%s", rec)
	}
}
