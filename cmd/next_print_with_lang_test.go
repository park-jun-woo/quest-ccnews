//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what printNext가 Lang이 있을 때 메타·본문·event6 규칙·submit 사용법과 "언어:" 줄을 모두 출력하는지 검증한다.

package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestPrintNext_WithLang(t *testing.T) {
	a := &session.Article{
		URL:   "https://example.com/a",
		Host:  "example.com",
		Lang:  "ko",
		Tries: 1,
	}
	var buf bytes.Buffer
	printNext(&buf, a, "the body text here")
	out := buf.String()

	for _, want := range []string{
		"다음 TODO 기사",
		"https://example.com/a",
		"example.com",
		"언어: ko",
		"the body text here",
		"event6 작성 규칙",
		"ccnews submit",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q\n---\n%s", want, out)
		}
	}
}
