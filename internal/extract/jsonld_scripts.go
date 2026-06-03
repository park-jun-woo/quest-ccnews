//ff:func feature=extract type=helper control=sequence level=error
//ff:what HTML을 파싱해 <script type="application/ld+json"> 블록들의 raw JSON 텍스트를 문서 순서대로 수집한다. 순수 함수(파싱 실패 시 nil).

package extract

import (
	"bytes"

	"golang.org/x/net/html"
)

// jsonLDScripts returns the raw inner text of every
// <script type="application/ld+json"> block, in document order. Pure — returns
// nil on parse error or when no such block exists.
func jsonLDScripts(htmlBytes []byte) []string {
	doc, err := html.Parse(bytes.NewReader(htmlBytes))
	if err != nil {
		return nil
	}
	var out []string
	collectLDScripts(doc, &out)
	return out
}
