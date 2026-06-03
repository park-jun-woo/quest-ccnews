//ff:func feature=extract type=helper control=sequence
//ff:what collectMeta가 트리의 <meta> 요소들에서 OG/article 필드를 누적 Fields에 반영하는지 검증한다.

package extract

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestCollectMeta(t *testing.T) {
	doc, err := html.Parse(strings.NewReader(`<html><head>` +
		`<meta property="og:title" content="OG Title">` +
		`<meta property="og:site_name" content="OG Site">` +
		`<meta name="article:author" content="Bob">` +
		`</head><body></body></html>`))
	if err != nil {
		t.Fatal(err)
	}
	var f Fields
	collectMeta(doc, &f)
	want := Fields{Title: "OG Title", MediaName: "OG Site", Author: "Bob"}
	if f != want {
		t.Fatalf("collectMeta got %+v, want %+v", f, want)
	}
}
