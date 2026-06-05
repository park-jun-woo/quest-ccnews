//ff:func feature=cli type=helper control=sequence
//ff:what runSubmit가 --event6 -로 stdin에서 event6를 읽어 정상 PASS 판정을 내는지 검증한다.

package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestRunSubmit_StdinEvent6(t *testing.T) {
	resetSubmitFlags(t)
	cache, file := writeCacheWarc(t, submitStructuredHTML(
		"Alice met Bob in Paris on Monday to sign the treaty. "+
			"The two leaders shook hands before the cameras and pledged to deepen cooperation "+
			"across trade, security and climate over the coming year, aides said afterward."))
	p := writeSessionWith(t, file)
	submitURL = "https://example.com/a"
	submitEvent6 = "-"
	sessionPath = p
	prev := submitCacheDir
	submitCacheDir = cache
	t.Cleanup(func() { submitCacheDir = prev })

	cmd := &cobra.Command{}
	cmd.SetIn(strings.NewReader(`{
		"who":{"value":"Alice","anchors":["Alice"]},
		"when":{"value":"Monday","anchors":["Monday"]},
		"what":{"value":"signed treaty","anchors":["sign the treaty"]}
	}`))
	var out bytes.Buffer
	cmd.SetOut(&out)
	if err := runSubmit(cmd, nil); err != nil {
		t.Fatalf("runSubmit stdin: %v", err)
	}
	if !strings.Contains(out.String(), "판정: PASS") {
		t.Errorf("output = %q", out.String())
	}
}
