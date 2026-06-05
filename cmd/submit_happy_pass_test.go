//ff:func feature=cli type=helper control=sequence
//ff:what runSubmit가 필수 앵커가 원문 substring인 event6에 대해 "판정: PASS"를 찍고 세션에 기사를 PASS로 잠가 저장하는지 검증한다.

package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/spf13/cobra"
)

func TestRunSubmit_HappyPass(t *testing.T) {
	resetSubmitFlags(t)
	cache, file := writeCacheWarc(t, submitStructuredHTML(
		"Alice met Bob in Paris on Monday to sign the treaty. "+
			"The two leaders shook hands before the cameras and pledged to deepen cooperation "+
			"across trade, security and climate over the coming year, aides said afterward."))
	p := writeSessionWith(t, file)
	evPath := writeEvent6File(t, `{
		"who":{"value":"Alice","anchors":["Alice"]},
		"when":{"value":"Monday","anchors":["Monday"]},
		"what":{"value":"signed treaty","anchors":["sign the treaty"]}
	}`)
	submitURL = "https://example.com/a"
	submitEvent6 = evPath
	sessionPath = p
	prev := submitCacheDir
	submitCacheDir = cache
	t.Cleanup(func() { submitCacheDir = prev })

	cmd := &cobra.Command{}
	var out bytes.Buffer
	cmd.SetOut(&out)
	if err := runSubmit(cmd, nil); err != nil {
		t.Fatalf("runSubmit: %v", err)
	}
	if !strings.Contains(out.String(), "판정: PASS") {
		t.Errorf("output = %q", out.String())
	}

	// Session should be persisted with the article locked to PASS.
	s2, err := session.Load(p)
	if err != nil {
		t.Fatal(err)
	}
	if s2.Articles[0].State != session.PASS {
		t.Errorf("article state = %s, want PASS", s2.Articles[0].State)
	}
}
