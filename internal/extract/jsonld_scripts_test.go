//ff:func feature=extract type=helper control=sequence
//ff:what jsonLDScripts가 ld+json 블록을 문서 순서대로 모으고 비-ld·빈 블록은 건너뛰며, 없으면 nil을 돌려주는지 검증한다.

package extract

import (
	"strings"
	"testing"
)

func TestJSONLDScripts(t *testing.T) {
	t.Run("collects in order, skips non-ld and empty", func(t *testing.T) {
		in := []byte(`<html><head>` +
			`<script type="application/ld+json">{"a":1}</script>` +
			`<script type="text/javascript">var x=1</script>` +
			`<script type="application/ld+json">   </script>` +
			`<script type="application/ld+json">{"b":2}</script>` +
			`</head><body></body></html>`)
		got := jsonLDScripts(in)
		if len(got) != 2 {
			t.Fatalf("want 2 scripts, got %d: %v", len(got), got)
		}
		if !strings.Contains(got[0], `"a":1`) || !strings.Contains(got[1], `"b":2`) {
			t.Fatalf("wrong order/content: %v", got)
		}
	})
	t.Run("none present", func(t *testing.T) {
		if got := jsonLDScripts([]byte(`<html><body>hi</body></html>`)); got != nil {
			t.Fatalf("want nil, got %v", got)
		}
	})
	t.Run("empty input parses to nil", func(t *testing.T) {
		// html.Parse never errors on a byte slice (the tokenizer is lenient),
		// so the err!=nil branch is unreachable; empty input still parses to a
		// document with no ld+json blocks and yields nil.
		if got := jsonLDScripts(nil); got != nil {
			t.Fatalf("want nil for empty input, got %v", got)
		}
	})
	t.Run("ld block in body, not just head", func(t *testing.T) {
		in := []byte(`<html><body>` +
			`<script type="application/ld+json">{"c":3}</script>` +
			`</body></html>`)
		got := jsonLDScripts(in)
		if len(got) != 1 || !strings.Contains(got[0], `"c":3`) {
			t.Fatalf("want one body-level ld block, got %v", got)
		}
	})
}
