//ff:func feature=cli type=helper control=sequence
//ff:what 테스트 헬퍼. 주어진 본문을 articleBody로 담은 JSON-LD NewsArticle HTML을 만든다. 신뢰 게이트(extract.Apply)를 통과시켜 앵커 게이트까지 도달하게 한다.

package cmd

import "strings"

// submitStructuredHTML wraps body as the articleBody of a minimal JSON-LD
// NewsArticle so extract.Apply's trust gate passes (structured source + title)
// and returns body as the anchor-target text. The body must be >= MinBodyLen
// and contain the event6 anchors the test submits.
func submitStructuredHTML(body string) string {
	esc := strings.ReplaceAll(body, `"`, `\"`)
	return `<html><head>` +
		`<script type="application/ld+json">` +
		`{"@type":"NewsArticle","headline":"Test headline","datePublished":"2026-06-04",` +
		`"articleBody":"` + esc + `"}` +
		`</script></head><body></body></html>`
}
