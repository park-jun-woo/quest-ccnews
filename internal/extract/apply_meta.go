//ff:func feature=extract type=helper control=selection
//ff:what 하나의 meta key/content 쌍을 대응하는 Fields 항목에 설정한다(필드별 첫 비어있지 않은 값 우선). og:title/og:site_name/article:published_time/article:author 매핑.

package extract

// applyMeta sets the matching Fields entry for one meta key/content pair (first
// non-empty wins per field).
func applyMeta(f *Fields, key, content string) {
	if content == "" {
		return
	}
	switch key {
	case "og:title":
		setIfEmpty(&f.Title, content)
	case "og:site_name":
		setIfEmpty(&f.MediaName, content)
	case "article:published_time":
		setIfEmpty(&f.PublishedAt, content)
	case "article:author":
		setIfEmpty(&f.Author, content)
	}
}
