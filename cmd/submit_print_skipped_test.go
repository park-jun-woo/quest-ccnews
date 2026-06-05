//ff:func feature=cli type=helper control=sequence
//ff:what printSubmitSkipped가 "판정: SKIPPED"·SkipReason·기사 상태를 찍고 앵커 게이트 미실행·재제출 불가 안내를 포함하는지 검증한다.

package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestPrintSubmitSkipped(t *testing.T) {
	a := &session.Article{State: session.SKIPPED, SkipReason: "no structured data"}
	var buf bytes.Buffer
	printSubmitSkipped(&buf, a)
	out := buf.String()
	if !strings.Contains(out, "판정: "+string(session.SKIPPED)) {
		t.Errorf("output = %q, want SKIPPED verdict line", out)
	}
	if !strings.Contains(out, "사유: no structured data") {
		t.Errorf("output = %q, want SkipReason line", out)
	}
	if !strings.Contains(out, "기사 상태: "+string(session.SKIPPED)) {
		t.Errorf("output = %q, want article state line", out)
	}
	if !strings.Contains(out, "재제출 불가") {
		t.Errorf("output = %q, want terminal/no-retry guidance", out)
	}
}
