//ff:func feature=session type=helper control=sequence
//ff:what 세션을 들여쓴 JSON으로 직렬화해 파일에 쓴다.

package session

import (
	"encoding/json"
	"os"
)

// Save: serialize the session as indented JSON and write it to path.
func (s *Session) Save(path string) error {
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0o644)
}
