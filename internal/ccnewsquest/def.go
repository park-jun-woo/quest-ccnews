//ff:func feature=gate type=helper control=sequence
//ff:what ccnews 퀘스트 정의(ccnewsDef)를 gate.Definition으로 반환해 cli.NewQuestCmd에 끼운다. userAgent·cacheDir는 WARC 재독 클라이언트 설정(Prepare가 사용; Phase013에서 Session.Meta 연결).

package ccnewsquest

import "github.com/park-jun-woo/reins/pkg/gate"

// Def returns the ccnews quest definition to wire into cli.NewQuestCmd. userAgent
// and cacheDir configure the WARC re-read client used by Prepare (Phase013 will
// source them from the session meta slot).
func Def(userAgent, cacheDir string) gate.Definition {
	return ccnewsDef{userAgent: userAgent, cacheDir: cacheDir}
}
