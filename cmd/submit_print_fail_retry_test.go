//ff:func feature=cli type=helper control=sequence
//ff:what printSubmit(FAIL, TODO)가 "판정: FAIL"과 "TODO 유지" 재시도 안내를 찍는지 검증한다.

package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/anchor"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestPrintSubmit_FailRetry(t *testing.T) {
	a := &session.Article{State: session.TODO, Tries: 1}
	var buf bytes.Buffer
	printSubmit(&buf, a, anchor.Result{Verdict: anchor.FAIL, Reason: "bad"})
	out := buf.String()
	if !strings.Contains(out, "판정: FAIL") {
		t.Errorf("output = %q", out)
	}
	if !strings.Contains(out, "TODO 유지") {
		t.Errorf("FAIL+TODO should mention retry, got %q", out)
	}
}
