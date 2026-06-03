//ff:func feature=extract type=helper control=sequence
//ff:what collectLDScripts가 트리에서 ld+json 블록 텍스트를 문서 순서로 모으고 비-ld·빈 블록은 건너뛰는지 검증한다.

package extract

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestCollectLDScripts(t *testing.T) {
	doc, err := html.Parse(strings.NewReader(`<html><head>` +
		`<script type="application/ld+json">{"a":1}</script>` +
		`<script type="text/javascript">var x=1</script>` +
		`<script type="application/ld+json">   </script>` +
		`<script type="application/ld+json">{"b":2}</script>` +
		`</head><body></body></html>`))
	if err != nil {
		t.Fatal(err)
	}
	var out []string
	collectLDScripts(doc, &out)
	if len(out) != 2 {
		t.Fatalf("want 2 scripts, got %d: %v", len(out), out)
	}
	if !strings.Contains(out[0], `"a":1`) || !strings.Contains(out[1], `"b":2`) {
		t.Fatalf("wrong order/content: %v", out)
	}
}
