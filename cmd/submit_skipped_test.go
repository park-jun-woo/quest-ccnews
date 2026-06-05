//ff:func feature=cli type=helper control=sequence
//ff:what runSubmit가 구조화 데이터 없는 기사에 대해 extract.Apply 신뢰 게이트에서 SKIPPED로 단락(앵커 게이트 미실행)하고, 세션을 SKIPPED로 잠그고 audit 레코드를 out으로 emit하는지 검증한다(Phase010-B).

package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/spf13/cobra"
)

func TestRunSubmit_SkippedShortCircuit(t *testing.T) {
	resetSubmitFlags(t)
	// No JSON-LD / OG / title → extract.Apply's trust gate fails → SKIPPED before
	// the anchor gate ever runs.
	cache, file := writeCacheWarc(t, "<html><body>no structured data here</body></html>")
	p := writeSessionWith(t, file)
	evPath := writeEvent6File(t, `{
		"who":{"value":"Alice","anchors":["Alice"]},
		"what":{"value":"signed treaty","anchors":["sign the treaty"]}
	}`)
	submitURL = "https://example.com/a"
	submitEvent6 = evPath
	sessionPath = p
	prev := submitCacheDir
	submitCacheDir = cache
	t.Cleanup(func() { submitCacheDir = prev })

	out := filepath.Join(t.TempDir(), "out.jsonl")
	prevOut := submitOut
	submitOut = out
	t.Cleanup(func() { submitOut = prevOut })

	cmd := &cobra.Command{}
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	if err := runSubmit(cmd, nil); err != nil {
		t.Fatalf("runSubmit: %v", err)
	}
	if !strings.Contains(buf.String(), "판정: SKIPPED") {
		t.Errorf("output = %q, want SKIPPED report", buf.String())
	}

	// Session persisted with the article locked to SKIPPED (anchor gate not reached
	// ⇒ no Tries increment, no verdict).
	s2, err := session.Load(p)
	if err != nil {
		t.Fatal(err)
	}
	a := s2.Articles[0]
	if a.State != session.SKIPPED {
		t.Errorf("article state = %s, want SKIPPED", a.State)
	}
	if a.Tries != 0 {
		t.Errorf("Tries = %d, want 0 (anchor gate must not run on SKIPPED)", a.Tries)
	}

	// The terminal SKIPPED article must have been swept to --out as an audit record.
	data, err := os.ReadFile(out)
	if err != nil {
		t.Fatalf("read out: %v", err)
	}
	if !strings.Contains(string(data), `"status":"SKIPPED"`) {
		t.Errorf("out = %q, want a SKIPPED audit record", string(data))
	}
}
