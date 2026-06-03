//ff:func feature=extract type=helper control=sequence
//ff:what Fields가 기사임을 사이트가 자기선언했는지 판정한다. 비어있지 않은 Title이 기사 마커이며 신뢰 게이트의 키. 순수 함수.

package extract

// hasArticle reports whether the source self-declared an Article at all. A
// non-empty Title is the marker: it is the field the trust gate keys on, and a
// site that declares no headline has not self-declared an article.
func (f Fields) hasArticle() bool { return f.Title != "" }
