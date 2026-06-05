//ff:func feature=gate type=helper control=sequence
//ff:what 규칙·동등성 테스트 공용 헬퍼 fld(value,앵커…)→*session.Field. testSource는 모든 테스트 앵커가 substring인 공용 원문(anchor.gateSource와 동일).

package ccnewsquest

import "github.com/park-jun-woo/quest-ccnews/internal/session"

// testSource mirrors anchor.gateSource: the body every test anchor is a substring of.
const testSource = "Alice met Bob in Paris on Monday to sign the treaty because peace mattered"

// fld builds a *session.Field from a value and its anchors.
func fld(value string, anchors ...string) *session.Field {
	return &session.Field{Value: value, Anchors: anchors}
}
