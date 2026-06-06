//ff:func feature=gate type=helper control=sequence level=error
//ff:what Render의 본문 덤프 단계를 추출한 헬퍼. Phase013 C의 "원문 본문(앵커 대상)" 블록을 Prepare와 동일한 공유 헬퍼 readArticleBody로 재독해 trimBody로 상한·trim 후 b에 기록한다. Render는 read-only이므로 본문 재독은 로컬 복사 Article(body)에 대해 하고 Meta를 쓰지 않는다. WARC 미가용·trust gate skip·디코드 실패는 best-effort로 무시(본문만 생략, 호출부 a.WARC != nil 가드 하에서만 진입). Render의 중첩깊이를 ≤2로 유지하기 위한 순수 구조 추출 — 출력 바이트는 추출 전과 동일.
package ccnewsquest

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

// renderBody appends the anchor-target body section to b, re-reading the article
// body through the same shared helper Prepare uses (readArticleBody) and capping it
// with trimBody. It is the extracted inner block of Render, kept byte-identical to
// the pre-extraction output; the caller guards entry with a.WARC != nil. Best-effort:
// a decode error or a WARC/trust-gate miss simply omits the body section.
func (d ccnewsDef) renderBody(s *quest.Session, it *quest.Item, b *strings.Builder) {
	userAgent, cacheDir := d.sourceConfig(s)
	var body session.Article
	if err := it.DecodePayload(&body); err != nil {
		return
	}
	bodyText, ok, err := readArticleBody(userAgent, cacheDir, &body)
	if err != nil || !ok {
		return
	}
	fmt.Fprintln(b, "=== 원문 본문 (앵커 대상) ===")
	fmt.Fprintln(b, trimBody(bodyText))
	fmt.Fprintln(b)
}
