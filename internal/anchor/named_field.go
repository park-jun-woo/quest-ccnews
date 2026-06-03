//ff:type feature=anchor type=model
//ff:what 필드 라벨과 event6 구조체 안의 Field 포인터를 묶는다. nil Field는 "부재"를 뜻한다. 게이트가 판정 사유에 어느 필드인지 붙이기 위해 쓴다.

package anchor

import "github.com/park-jun-woo/quest-ccnews/internal/session"

// namedField pairs a field label with the pointer into the event6 struct so the
// gate can report which field a verdict concerns. A nil Field means "absent".
type namedField struct {
	name string
	f    *session.Field
}
