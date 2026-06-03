//ff:func feature=extract type=helper control=sequence
//ff:what collectText가 body 텍스트(Hello/World)는 모으고 script 본문과 head/title은 누락시키는지 검증한다.

package extract

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestCollectText(t *testing.T) {
	doc, err := html.Parse(strings.NewReader(`<html><head><title>T</title></head><body><p>Hello</p><script>var x=1;</script><p>World</p></body></html>`))
	if err != nil {
		t.Fatal(err)
	}
	var b strings.Builder
	collectText(doc, &b)
	got := strings.Join(strings.Fields(b.String()), " ")
	if !strings.Contains(got, "Hello") || !strings.Contains(got, "World") {
		t.Fatalf("missing body text: %q", got)
	}
	if strings.Contains(got, "var x") {
		t.Fatalf("script leaked: %q", got)
	}
	if strings.Contains(got, "T") {
		t.Fatalf("head/title leaked: %q", got)
	}
}
