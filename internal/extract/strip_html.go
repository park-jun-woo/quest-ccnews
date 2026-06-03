//ff:func feature=extract type=helper control=sequence level=error
//ff:what HTML 바이트에서 script/style/head 등을 빼고 모든 텍스트 노드를 모아 공백 정규화한 전문 텍스트를 만든다. 순수 함수(파싱 실패 시 빈 문자열).

package extract

import (
	"bytes"
	"strings"

	"golang.org/x/net/html"
)

// stripHTML parses the HTML and returns its visible text with <script>, <style>,
// <noscript>, <template> and <head> subtrees removed and whitespace collapsed to
// single spaces. This is the anchor-target fallback when JSON-LD declares no
// articleBody. Pure — on a parse error it returns "" (no usable text).
func stripHTML(htmlBytes []byte) string {
	doc, err := html.Parse(bytes.NewReader(htmlBytes))
	if err != nil {
		return ""
	}
	var b strings.Builder
	collectText(doc, &b)
	return strings.Join(strings.Fields(b.String()), " ")
}
