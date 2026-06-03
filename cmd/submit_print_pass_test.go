//ff:func feature=cli type=helper control=sequence
//ff:what printSubmit(PASS)가 "판정: PASS"·"사유: ok"를 찍고 재시도 안내를 포함하지 않는지 검증한다.

package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/anchor"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestPrintSubmit_Pass(t *testing.T) {
	a := &session.Article{State: session.PASS}
	var buf bytes.Buffer
	printSubmit(&buf, a, anchor.Result{Verdict: anchor.PASS, Reason: "ok"})
	out := buf.String()
	if !strings.Contains(out, "판정: PASS") || !strings.Contains(out, "사유: ok") {
		t.Errorf("output = %q", out)
	}
	if strings.Contains(out, "재제출") {
		t.Error("PASS output should not include retry guidance")
	}
}
