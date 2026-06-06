//ff:func feature=gate type=helper control=sequence
//ff:what ccnews 퀘스트 정의(ccnewsDef)를 gate.Definition으로 반환해 cli.NewQuestCmd에 끼운다. userAgent·cacheDir는 WARC 재독 기본값(하위호환 fallback; Prepare/Render는 Session.Meta에서 우선 소싱). pick-time robots 평가 캐시를 1회 생성해 주입한다(호스트당 1회 fetch — Phase013 A).

package ccnewsquest

import "github.com/park-jun-woo/reins/pkg/gate"

// Def returns the ccnews quest definition to wire into cli.NewQuestCmd. userAgent
// and cacheDir are the WARC re-read defaults; Prepare/Render prefer the values in
// quest.Session.Meta and fall back to these (Phase013 B). A fresh pick-time robots
// cache is created here so every Prepare in the process shares one cache and each
// host's robots.txt is fetched at most once (Phase013 A).
func Def(userAgent, cacheDir string) gate.Definition {
	return ccnewsDef{userAgent: userAgent, cacheDir: cacheDir, robots: newRobotsCache()}
}
