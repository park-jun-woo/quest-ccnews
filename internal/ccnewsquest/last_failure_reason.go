//ff:func feature=gate type=helper control=sequence level=error
//ff:what lastFailureReason(it) — 아이템 로그의 마지막 항목 Reason을 반환. 로그가 비었으면 "". Render의 직전-실패 게이팅 삼중 if를 깊이 ≤2로 낮추기 위한 순수 구조 추출 — 반환값은 추출 전 it.Log[n-1].Reason과 동일.
package ccnewsquest

import "github.com/park-jun-woo/reins/pkg/quest"

// lastFailureReason returns the Reason of the item's last log entry, or "" when the
// log is empty. It is the extracted tail-read of Render's last-failure gating, kept
// behavior-identical to reading it.Log[len(it.Log)-1].Reason directly.
func lastFailureReason(it *quest.Item) string {
	if n := len(it.Log); n > 0 {
		return it.Log[n-1].Reason
	}
	return ""
}
