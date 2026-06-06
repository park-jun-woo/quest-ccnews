//ff:func feature=gate type=helper control=sequence level=error
//ff:what Prepare·Render 공유 본문 헬퍼(Phase013 C). userAgent·cacheDir로 ingest.Client를 만들어 a.WARC를 ReadBody로 재독하고 extract.Apply로 추출+신뢰 게이트를 돌린다. (bodyText, ok, err)을 돌려준다 — bodyText는 게이트 입력(Source)과 동일한 앵커 대상 텍스트. a를 제자리 변이(PASS=Extracted/Lang, SKIP=State/SkipReason)하므로 호출자가 영속 여부를 정한다(Prepare는 SetPayload, Render는 읽기만). WARC 디스크 IO는 결정론적(로컬 캐시, 라이브 fetch 아님).
package ccnewsquest

import (
	"fmt"

	"github.com/park-jun-woo/quest-ccnews/internal/extract"
	"github.com/park-jun-woo/quest-ccnews/internal/ingest"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// readArticleBody re-reads the article body from its WARC locator (a deterministic
// local-cache disk read, not a live fetch) and runs extract.Apply (structured-field
// extraction + trust gate). It is the single shared text path so `next` (Render) and
// `submit` (Prepare) always see the byte-identical anchor-target body — the same
// string Prepare hands the gate as Context.Source — avoiding surface-form drift that
// would cause false FAILs (Phase013 C). It mutates a in place (PASS → Extracted/Lang,
// trust FAIL → State=SKIPPED + SkipReason); callers decide whether to persist that
// (Prepare SetPayloads it; Render is read-only and discards). bodyText is empty when
// ok is false.
func readArticleBody(userAgent, cacheDir string, a *session.Article) (bodyText string, ok bool, err error) {
	client := ingest.NewClient(userAgent, cacheDir)
	htmlBytes, err := client.ReadBody(a.WARC)
	if err != nil {
		return "", false, fmt.Errorf("원문 재독 실패 (%s): %w", a.URL, err)
	}
	bodyText, ok = extract.Apply(a, htmlBytes)
	return bodyText, ok, nil
}
