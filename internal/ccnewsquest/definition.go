//ff:type feature=gate type=model
//ff:what ccnews의 gate.Definition 구현체(ccnewsDef). 판정은 Rules() 카탈로그(앵커 규칙 6개)를 reins gate.Evaluate가 레벨집계(필수 FAIL>선택 REVIEW)해서 한다 — ccnews는 진짜 선후가드가 없어 그래프 불요. userAgent·cacheDir는 WARC 재독 기본값(하위호환 fallback)이며 Phase013부터 Session.Meta(G2: cache_dir/user_agent)에서 우선 소싱한다. robots는 pick-time robots 평가 캐시 포인터로, 한 프로세스 동안 호스트당 1회만 fetch하게 공유된다(Phase013 A).

package ccnewsquest

// ccnewsDef is ccnews's gate.Definition implementation. The verdict is computed
// from the Rules() catalog (the six anchor rules) via reins gate.Evaluate's level
// aggregation: a fired required-* Fail rule beats a present-but-unverifiable
// optional Review (Level Fail > Review), which preserves the original anchor.Gate
// ordering verdict-for-verdict. ccnews has no genuine precedence guards, so no
// defeat graph (Evaluator) is needed — the flat Rules() path suffices.
//
// userAgent and cacheDir are the WARC re-read defaults: Prepare/Render source them
// from quest.Session.Meta (G2: cache_dir/user_agent) when present, falling back to
// these receiver values for sessions that predate Phase013 (backward compatible).
// robots is the shared pick-time robots evaluation cache (a pointer so the by-value
// receiver shares one cache across the process), making each host's robots.txt
// fetched at most once per run (Phase013 A).
type ccnewsDef struct {
	userAgent string
	cacheDir  string
	robots    *robotsCache
}
