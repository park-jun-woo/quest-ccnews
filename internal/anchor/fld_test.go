//ff:func feature=anchor type=helper control=sequence
//ff:what 게이트 테스트 헬퍼. value와 앵커들로 session.Field 포인터를 만든다. gateSource는 모든 테스트 앵커가 substring인 공용 원문.

package anchor

import "github.com/park-jun-woo/quest-ccnews/internal/session"

// gateSource is the source body that every gate-test anchor is a substring of
// (after normalize).
const gateSource = "Alice met Bob in Paris on Monday to sign the treaty because peace mattered"

// fld builds a *session.Field from a value and its anchors.
func fld(value string, anchors ...string) *session.Field {
	return &session.Field{Value: value, Anchors: anchors}
}
