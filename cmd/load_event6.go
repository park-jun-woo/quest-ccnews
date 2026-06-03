//ff:func feature=cli type=helper control=sequence level=error
//ff:what submit의 event6 입력 로더. "-"면 stdin, 아니면 파일에서 JSON을 읽어 session.Event6로 역직렬화한다.

package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// loadEvent6 reads an event6 JSON document from path ("-" means stdin) and
// deserializes it into a session.Event6 (value=English, anchors=original-language
// surface forms). Thin IO + parse.
func loadEvent6(path string, stdin io.Reader) (*session.Event6, error) {
	var b []byte
	var err error
	if path == "-" {
		b, err = io.ReadAll(stdin)
	} else {
		b, err = os.ReadFile(path)
	}
	if err != nil {
		return nil, err
	}
	var ev session.Event6
	if err := json.Unmarshal(b, &ev); err != nil {
		return nil, fmt.Errorf("parse event6 JSON %s: %w", path, err)
	}
	return &ev, nil
}
