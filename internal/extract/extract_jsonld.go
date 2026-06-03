//ff:func feature=extract type=helper control=iteration dimension=1 level=error
//ff:what JSON-LD 블록들을 파싱해 평탄화하고, 첫 Article 노드를 Fields로 매핑해 돌려준다. 깨진 블록은 개별 무시. 순수 함수.

package extract

import "encoding/json"

// extractJSONLD parses each ld+json block, flattens @graph/arrays/nesting into a
// flat list of objects, and maps the first Article-typed object to Fields. It
// reports ok=false when no Article object is found. Pure — malformed blocks are
// skipped individually.
func extractJSONLD(scripts []string) (Fields, bool) {
	for _, s := range scripts {
		var raw any
		if json.Unmarshal([]byte(s), &raw) != nil {
			continue
		}
		var objs []map[string]any
		flattenLD(raw, &objs)
		if f, ok := firstArticle(objs); ok {
			return f, true
		}
	}
	return Fields{}, false
}
