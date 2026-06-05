//ff:type feature=gate type=model
//ff:what ccnews의 gate.Definition 구현체(ccnewsDef). 판정은 Rules() 카탈로그(앵커 규칙 6개)를 reins gate.Evaluate가 레벨집계(필수 FAIL>선택 REVIEW)해서 한다 — ccnews는 진짜 선후가드가 없어 그래프 불요. userAgent·cacheDir는 소비자 설정(Prepare가 WARC 재독에 쓰는 ingest.Client 인자) — Phase013에서 Session.Meta(G2)로 공급된다.

package ccnewsquest

// ccnewsDef is ccnews's gate.Definition implementation. The verdict is computed
// from the Rules() catalog (the six anchor rules) via reins gate.Evaluate's level
// aggregation: a fired required-* Fail rule beats a present-but-unverifiable
// optional Review (Level Fail > Review), which preserves the original anchor.Gate
// ordering verdict-for-verdict. ccnews has no genuine precedence guards, so no
// defeat graph (Evaluator) is needed — the flat Rules() path suffices.
//
// userAgent and cacheDir are consumer configuration: Prepare hands them to
// ingest.NewClient to re-read the article body from its WARC locator. In Phase013
// they will be sourced from quest.Session.Meta (G2); for now they are set by Def's
// caller.
type ccnewsDef struct {
	userAgent string
	cacheDir  string
}
