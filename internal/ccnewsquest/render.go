//ff:func feature=gate type=helper control=sequence level=error
//ff:what Definition.Render(s, it). 다음 TODO 기사에 대한 event6 작성 프롬프트를 만든다. value=영어, anchors=원문 표면형, 필수 who/what, 선택 when/where/how/why(앵커 0개면 REVIEW). reins submit은 JSON stdin/파일 제출이므로 `ccnews submit --key <URL> --in -` 사용법과 스키마를 명시한다. 직전 실패 사유는 아이템 로그 마지막 항목에서 읽는다. Phase013 C: cacheDir·UA를 s.Meta에서 소싱(미기록 시 리시버 fallback)해 공유 헬퍼 readArticleBody로 WARC를 재독, Prepare의 Source와 바이트 동일한 앵커 대상 본문을 길이 상한·trim하여 출력에 포함한다(에이전트가 next만으로 event6 작성). Render는 read-only — 본문 재독은 로컬 복사 Article에 대해 하고 Meta를 쓰지 않는다.
package ccnewsquest

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

// renderBodyMax caps the body text dumped into the next prompt so a very long
// article does not blow up the prompt; the agent still sees enough to pick surface
// anchors, and submit re-reads the full body for verification.
const renderBodyMax = 8000

// Render returns the agent's event6 authoring prompt for the next TODO article. It
// states the value=English / anchors=original-surface-form contract, the required
// who/what vs optional when/where/how/why split, and the reins submit invocation
// `ccnews submit --key <URL> --in -`. The last failure reason (if any) is read from
// the item's log tail so a retrying agent sees why it failed.
//
// Phase013 C: it re-reads the article body through the same shared helper Prepare
// uses (readArticleBody), sourcing cacheDir/UA from the session Meta (Phase013 B),
// and includes that byte-identical anchor-target text (capped to renderBodyMax) so
// the agent can author event6 from `next` alone. Render is read-only: it re-reads
// into a throwaway Article copy and never writes to Meta (next does not Save).
func (d ccnewsDef) Render(s *quest.Session, it *quest.Item) (string, error) {
	var b strings.Builder

	url := it.Key
	host, lang := "", ""
	var a session.Article
	if err := it.DecodePayload(&a); err == nil {
		host, lang = a.Host, a.Lang
	}

	fmt.Fprintln(&b, "=== 다음 TODO 기사 ===")
	fmt.Fprintf(&b, "URL : %s\n", url)
	if host != "" {
		fmt.Fprintf(&b, "호스트: %s\n", host)
	}
	if lang != "" {
		fmt.Fprintf(&b, "언어: %s\n", lang)
	}
	fmt.Fprintf(&b, "시도: %d/%d\n", it.Tries, session.MaxTries)

	if n := len(it.Log); n > 0 {
		if r := it.Log[n-1].Reason; r != "" {
			fmt.Fprintf(&b, "직전 실패: %s\n", r)
		}
	}
	fmt.Fprintln(&b)

	// C. Include the anchor-target body so the agent can author event6 from `next`
	// alone. readArticleBody mutates its Article argument; a is a throwaway decode
	// here (Render persists nothing), so the mutation is harmless. Best-effort: if
	// the WARC is unavailable or the trust gate would skip, the body is simply
	// omitted and the authoring rules below still print.
	if a.WARC != nil {
		d.renderBody(s, it, &b)
	}

	fmt.Fprintln(&b, "=== event6 작성 규칙 ===")
	fmt.Fprintln(&b, "기사 원문에서 육하원칙(event6)을 만들어 제출하세요.")
	fmt.Fprintln(&b, "  - value  : 영어로 산출 (날짜는 ISO, 숫자는 정규화)")
	fmt.Fprintln(&b, "  - anchors: 원문에 글자 그대로 나타난 표면형 토큰들 (인명·날짜·숫자·지명)")
	fmt.Fprintln(&b, "  - 필수: who·what (값+앵커 필수, 비우면 FAIL)")
	fmt.Fprintln(&b, "  - 선택: when·where·how·why (있으면 앵커, 앵커 0개면 REVIEW)")
	fmt.Fprintln(&b, "  - 게이트는 anchors가 원문 substring인지만 본다. value는 검증 대상 아님.")
	fmt.Fprintln(&b, "  - 앵커가 하나라도 원문에 없으면 FAIL(환각). tries 소진 시 DONE.")
	fmt.Fprintln(&b)

	fmt.Fprintln(&b, "=== 제출 ===")
	fmt.Fprintf(&b, "  ccnews submit --key %q --in -\n", url)
	fmt.Fprintln(&b, "  event6 JSON 예: {\"who\":{\"value\":\"...\",\"anchors\":[\"...\"]},"+
		"\"when\":{...},\"what\":{...},\"where\":null,\"how\":null,\"why\":null}")

	return b.String(), nil
}

// trimBody caps body text to renderBodyMax bytes, appending an ellipsis marker when
// truncated so the agent knows the full body is longer (submit re-reads all of it).
func trimBody(s string) string {
	if len(s) <= renderBodyMax {
		return s
	}
	return s[:renderBodyMax] + "\n…(본문 일부 생략; submit은 전체 본문으로 검증)"
}
