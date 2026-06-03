//ff:func feature=extract type=helper control=sequence level=error
//ff:what HTML을 파싱해 <meta> 태그에서 OG/article 폴백 필드를 수집하고, 최소 title이 있으면 ok=true. 순수 함수(파싱 실패 또는 og:title 없으면 ok=false).

package extract

import (
	"bytes"

	"golang.org/x/net/html"
)

// extractMeta scans <meta> tags for the OG/article fallback fields and reports
// ok=true when at least a title was found (the site self-declared an article via
// OpenGraph). Pure — returns ok=false on parse error or when no og:title exists.
func extractMeta(htmlBytes []byte) (Fields, bool) {
	doc, err := html.Parse(bytes.NewReader(htmlBytes))
	if err != nil {
		return Fields{}, false
	}
	var f Fields
	collectMeta(doc, &f)
	return f, f.Title != ""
}
