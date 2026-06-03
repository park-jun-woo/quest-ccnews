//ff:func feature=extract type=helper control=sequence
//ff:what extractJSONLD가 평이/@graph/배열/타입배열에서 기사를 찾고, 깨진 블록은 건너뛰고 뒷블록을 쓰며, 비기사·빈입력은 ok=false임을 검증한다.

package extract

import "testing"

func TestExtractJSONLD(t *testing.T) {
	t.Run("plain NewsArticle", func(t *testing.T) {
		s := `{"@type":"NewsArticle","headline":"H","articleBody":"B"}`
		f, ok := extractJSONLD([]string{s})
		if !ok || f.Title != "H" || f.Body != "B" {
			t.Fatalf("got %+v ok=%v", f, ok)
		}
	})
	t.Run("@graph", func(t *testing.T) {
		s := `{"@graph":[{"@type":"WebPage"},{"@type":"Article","headline":"G"}]}`
		f, ok := extractJSONLD([]string{s})
		if !ok || f.Title != "G" {
			t.Fatalf("got %+v ok=%v", f, ok)
		}
	})
	t.Run("array of nodes", func(t *testing.T) {
		s := `[{"@type":"Person","name":"x"},{"@type":"NewsArticle","headline":"AR"}]`
		f, ok := extractJSONLD([]string{s})
		if !ok || f.Title != "AR" {
			t.Fatalf("got %+v ok=%v", f, ok)
		}
	})
	t.Run("type as array", func(t *testing.T) {
		s := `{"@type":["WebPage","NewsArticle"],"headline":"TA"}`
		f, ok := extractJSONLD([]string{s})
		if !ok || f.Title != "TA" {
			t.Fatalf("got %+v ok=%v", f, ok)
		}
	})
	t.Run("malformed block skipped, later block used", func(t *testing.T) {
		bad := `{not json`
		good := `{"@type":"Article","headline":"OK"}`
		f, ok := extractJSONLD([]string{bad, good})
		if !ok || f.Title != "OK" {
			t.Fatalf("got %+v ok=%v", f, ok)
		}
	})
	t.Run("no article object", func(t *testing.T) {
		s := `{"@type":"WebPage","name":"home"}`
		_, ok := extractJSONLD([]string{s})
		if ok {
			t.Fatalf("expected ok=false for non-article")
		}
	})
	t.Run("empty input", func(t *testing.T) {
		_, ok := extractJSONLD(nil)
		if ok {
			t.Fatalf("expected ok=false for nil scripts")
		}
	})
}
