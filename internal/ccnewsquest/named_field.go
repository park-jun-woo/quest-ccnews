//ff:type feature=gate type=model
//ff:what 필드 라벨과 event6 안의 Field 포인터를 묶는다(anchor.namedField 이식). nil Field는 "부재". 규칙이 발동 Fact에 어느 필드인지 붙이는 데 쓴다.

package ccnewsquest

import "github.com/park-jun-woo/quest-ccnews/internal/session"

// namedField pairs a field label with the pointer into the event6 struct so a rule
// can report which field its Fact concerns. A nil Field means "absent". Ported from
// anchor.namedField.
type namedField struct {
	name string
	f    *session.Field
}
