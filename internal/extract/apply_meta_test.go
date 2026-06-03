//ff:func feature=extract type=helper control=sequence
//ff:what applyMeta가 모든 키를 매핑하고(필드별 첫 값 우선), 빈 content와 미지의 키는 무시하는지 검증한다.

package extract

import "testing"

func TestApplyMeta(t *testing.T) {
	t.Run("maps all keys, first wins", func(t *testing.T) {
		var f Fields
		applyMeta(&f, "og:title", "T")
		applyMeta(&f, "og:title", "T2")
		applyMeta(&f, "og:site_name", "Site")
		applyMeta(&f, "article:published_time", "2026-01-01")
		applyMeta(&f, "article:author", "Auth")
		want := Fields{Title: "T", MediaName: "Site", PublishedAt: "2026-01-01", Author: "Auth"}
		if f != want {
			t.Fatalf("got %+v want %+v", f, want)
		}
	})
	t.Run("empty content ignored", func(t *testing.T) {
		var f Fields
		applyMeta(&f, "og:title", "")
		if f.Title != "" {
			t.Fatalf("empty content set title: %q", f.Title)
		}
	})
	t.Run("unknown key ignored", func(t *testing.T) {
		var f Fields
		applyMeta(&f, "og:image", "img.png")
		if f != (Fields{}) {
			t.Fatalf("unknown key mutated fields: %+v", f)
		}
	})
}
