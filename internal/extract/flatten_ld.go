//ff:func feature=extract type=helper control=selection
//ff:what 디코딩된 JSON-LD 값을 순회하며 발견한 모든 객체를 objs에 추가한다(배열/@graph 하강). extractJSONLD의 재귀 헬퍼.

package extract

// flattenLD walks a decoded JSON-LD value, appending every object it finds
// (descending into arrays and @graph) to objs. Recursive helper of extractJSONLD.
func flattenLD(v any, objs *[]map[string]any) {
	switch t := v.(type) {
	case []any:
		for _, e := range t {
			flattenLD(e, objs)
		}
	case map[string]any:
		*objs = append(*objs, t)
		if g, ok := t["@graph"]; ok {
			flattenLD(g, objs)
		}
	}
}
