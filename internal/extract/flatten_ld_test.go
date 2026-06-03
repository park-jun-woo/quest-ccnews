//ff:func feature=extract type=helper control=sequence
//ff:what flattenLD가 평객체/객체배열/@graph중첩을 모두 평탄화하고 스칼라는 무시하는지 검증한다.

package extract

import "testing"

func TestFlattenLD(t *testing.T) {
	t.Run("plain object", func(t *testing.T) {
		var objs []map[string]any
		flattenLD(map[string]any{"@type": "Article"}, &objs)
		if len(objs) != 1 {
			t.Fatalf("want 1 obj, got %d", len(objs))
		}
	})
	t.Run("array of objects", func(t *testing.T) {
		var objs []map[string]any
		flattenLD([]any{
			map[string]any{"@type": "A"},
			map[string]any{"@type": "B"},
		}, &objs)
		if len(objs) != 2 {
			t.Fatalf("want 2 objs, got %d", len(objs))
		}
	})
	t.Run("@graph nesting", func(t *testing.T) {
		var objs []map[string]any
		flattenLD(map[string]any{
			"@graph": []any{
				map[string]any{"@type": "WebPage"},
				map[string]any{"@type": "NewsArticle"},
			},
		}, &objs)
		// container object + 2 graph members
		if len(objs) != 3 {
			t.Fatalf("want 3 objs, got %d", len(objs))
		}
	})
	t.Run("scalar ignored", func(t *testing.T) {
		var objs []map[string]any
		flattenLD("string", &objs)
		flattenLD(42.0, &objs)
		flattenLD(nil, &objs)
		if len(objs) != 0 {
			t.Fatalf("want 0 objs, got %d", len(objs))
		}
	})
}
