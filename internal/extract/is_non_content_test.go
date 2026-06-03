//ff:func feature=extract type=helper control=iteration dimension=1
//ff:what isNonContent가 비콘텐츠 요소(script/style/noscript/template/head)는 true, 콘텐츠 요소(p/div/body/span)는 false로 판정하는지 검증한다.

package extract

import (
	"testing"

	"golang.org/x/net/html/atom"
)

func TestIsNonContent(t *testing.T) {
	for _, a := range []atom.Atom{atom.Script, atom.Style, atom.Noscript, atom.Template, atom.Head} {
		if !isNonContent(a) {
			t.Fatalf("isNonContent(%v) = false, want true", a)
		}
	}
	for _, a := range []atom.Atom{atom.P, atom.Div, atom.Body, atom.Span} {
		if isNonContent(a) {
			t.Fatalf("isNonContent(%v) = true, want false", a)
		}
	}
}
