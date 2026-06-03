//ff:func feature=cli type=helper control=sequence
//ff:what printSubmit(FAIL, DONE)가 "DONE으로 잠겼습니다" 안내를 찍는지 검증한다.

package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/anchor"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestPrintSubmit_FailDone(t *testing.T) {
	a := &session.Article{State: session.DONE, Tries: session.MaxTries}
	var buf bytes.Buffer
	printSubmit(&buf, a, anchor.Result{Verdict: anchor.FAIL, Reason: "final"})
	out := buf.String()
	if !strings.Contains(out, "DONE으로 잠겼습니다") {
		t.Errorf("FAIL+DONE should mention lock, got %q", out)
	}
}
