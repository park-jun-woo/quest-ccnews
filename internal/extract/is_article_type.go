//ff:func feature=extract type=helper control=selection
//ff:what JSON-LD @type 값(문자열 또는 문자열 배열)이 허용된 기사 타입을 가리키는지 판정한다. 순수 함수.

package extract

// articleTypes is the set of schema.org @type values accepted as a news article.
var articleTypes = map[string]bool{
	"NewsArticle":           true,
	"Article":               true,
	"ReportageNewsArticle":  true,
	"BackgroundNewsArticle": true,
	"OpinionNewsArticle":    true,
	"AnalysisNewsArticle":   true,
}

// isArticleType reports whether a JSON-LD @type value (string or array of
// strings) names an accepted article type.
func isArticleType(v any) bool {
	switch t := v.(type) {
	case string:
		return articleTypes[t]
	case []any:
		return anyArticleType(t)
	}
	return false
}
