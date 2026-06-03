//ff:func feature=cli type=helper control=sequence
//ff:what printNext가 Lang이 비면 "언어:" 줄을 출력하지 않는지 검증한다.

package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestPrintNext_NoLang(t *testing.T) {
	a := &session.Article{URL: "https://example.com/b", Host: "example.com"}
	var buf bytes.Buffer
	printNext(&buf, a, "body")
	if strings.Contains(buf.String(), "언어:") {
		t.Error("should not print language line when Lang empty")
	}
}
