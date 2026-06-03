//ff:func feature=extract type=helper control=iteration dimension=1
//ff:what isLDJSON이 정확한 타입/대문자/파라미터부착은 true, 다른 타입/타입속성없음은 false로 판정하는지 테이블로 검증한다.

package extract

import (
	"testing"

	"golang.org/x/net/html"
)

func TestIsLDJSON(t *testing.T) {
	cases := []struct {
		name  string
		attrs []html.Attribute
		want  bool
	}{
		{"exact", []html.Attribute{{Key: "type", Val: "application/ld+json"}}, true},
		{"uppercase key/val", []html.Attribute{{Key: "TYPE", Val: "APPLICATION/LD+JSON"}}, true},
		{"with params", []html.Attribute{{Key: "type", Val: " application/ld+json; charset=utf-8 "}}, true},
		{"other type", []html.Attribute{{Key: "type", Val: "text/javascript"}}, false},
		{"no type attr", []html.Attribute{{Key: "id", Val: "x"}}, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := isLDJSON(nodeWithAttrs(c.attrs...)); got != c.want {
				t.Fatalf("isLDJSON = %v, want %v", got, c.want)
			}
		})
	}
}
