//ff:func feature=extract type=helper control=iteration dimension=1
//ff:what 평탄화된 JSON-LD 객체 목록에서 첫 Article 타입 객체를 찾아 Fields로 매핑한다. 없으면 ok=false. 순수 함수.

package extract

// firstArticle scans a flattened list of JSON-LD objects and maps the first
// Article-typed object to Fields. It reports ok=false when none is an article.
// Pure.
func firstArticle(objs []map[string]any) (Fields, bool) {
	for _, o := range objs {
		if isArticleType(o["@type"]) {
			return mapArticle(o), true
		}
	}
	return Fields{}, false
}
