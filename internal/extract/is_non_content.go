//ff:func feature=extract type=helper control=selection
//ff:what 요소 서브트리가 보이는 기사 텍스트를 담지 않아 앵커 대상에서 제외해야 하는지 판정한다(script/style/noscript/template/head).

package extract

import "golang.org/x/net/html/atom"

// isNonContent reports whether an element subtree carries no visible article
// text and must be dropped from the anchor target.
func isNonContent(a atom.Atom) bool {
	switch a {
	case atom.Script, atom.Style, atom.Noscript, atom.Template, atom.Head:
		return true
	}
	return false
}
