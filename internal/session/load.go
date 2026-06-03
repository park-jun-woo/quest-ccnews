//ff:func feature=session type=helper control=sequence
//ff:what 세션 JSON 파일을 읽어 Session으로 역직렬화한다. nil Hosts 맵은 빈 맵으로 초기화한다.

package session

import (
	"encoding/json"
	"fmt"
	"os"
)

// Load: read a session JSON file and deserialize it into a Session.
func Load(path string) (*Session, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var s Session
	if err := json.Unmarshal(b, &s); err != nil {
		return nil, fmt.Errorf("parse session file %s: %w", path, err)
	}
	if s.Hosts == nil {
		s.Hosts = map[string]*Host{}
	}
	return &s, nil
}
