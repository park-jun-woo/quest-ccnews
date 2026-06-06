//ff:func feature=event6 type=helper control=sequence
//ff:what decodeEvent6 — 모델 출력에서 event6를 관용 디코드. 펜스 제거→첫 균형 JSON 객체 추출→Unmarshal. 성공 시 (ev,true), 실패 시 (zero,false). 포맷 잡음을 재시도 가능한 FAIL로 흡수하기 위한 입력단.

package ccnewsquest

import (
	"encoding/json"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// decodeEvent6 leniently decodes a model's raw submission into an Event6. It first
// strips any wrapping markdown code fence, then scans for the first balanced JSON
// object so stray prose around the object does not defeat the decode, and finally
// Unmarshals that slice. It returns (ev, true) on success and (zero, false) on any
// failure — Prepare turns the false case into a retryable FAIL verdict rather than a
// Go error, so a single bit of format noise from a small local model does not abort
// the unattended agent loop.
func decodeEvent6(raw []byte) (session.Event6, bool) {
	var zero session.Event6

	stripped := stripMarkdownFences(string(raw))
	obj, ok := firstJSONObject(stripped)
	if !ok {
		return zero, false
	}

	var ev session.Event6
	if err := json.Unmarshal([]byte(obj), &ev); err != nil {
		return zero, false
	}
	return ev, true
}
