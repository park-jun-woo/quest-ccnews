//ff:func feature=ingestion type=helper control=sequence level=error
//ff:what reins Session.Meta(map[string]any)에서 키 하나를 꺼내 JSON 왕복으로 타입 dst에 디코드한다. Load 직후 Meta 값은 map[string]any이므로 재마샬→언마샬해 session.Ingestion 등 구조체로 되돌린다. 키 부재면 (false,nil), 존재하면 디코드 후 (true,err).

package runcmd

import (
	"encoding/json"

	"github.com/park-jun-woo/reins/pkg/quest"
)

// decodeMeta extracts the value stored under key in the reins session's Meta slot
// and decodes it into dst via a JSON round-trip. After quest.Load a Meta value is
// a generic map[string]any (JSON's default), so re-marshalling and unmarshalling is
// how we recover a typed struct like session.Ingestion. It reports ok=false when
// the key is absent (dst untouched); otherwise it decodes and returns any error.
func decodeMeta(s *quest.Session, key string, dst any) (ok bool, err error) {
	v, present := s.GetMeta(key)
	if !present {
		return false, nil
	}
	b, err := json.Marshal(v)
	if err != nil {
		return true, err
	}
	if err := json.Unmarshal(b, dst); err != nil {
		return true, err
	}
	return true, nil
}
