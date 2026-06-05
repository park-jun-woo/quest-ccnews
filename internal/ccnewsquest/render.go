//ff:func feature=gate type=helper control=sequence level=error
//ff:what Definition.Render. 다음 TODO 기사에 대한 event6 작성 프롬프트를 만든다(cmd/next_print.go 이식). value=영어, anchors=원문 표면형, 필수 who/what, 선택 when/where/how/why(앵커 0개면 REVIEW). reins submit은 JSON stdin/파일 제출이므로 `ccnews submit --key <URL> --in -` 사용법과 스키마를 명시한다. 직전 실패 사유는 아이템 로그 마지막 항목에서 읽는다. (본문 텍스트는 Prepare가 WARC 재독으로 공급하므로 Render는 작성 규칙만 — Phase013 run 배선 전까지.)

package ccnewsquest

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

// Render returns the agent's event6 authoring prompt for the next TODO article
// (ported from cmd/next_print.go's printNext, minus the body dump: under reins the
// anchor-target body is re-read by Prepare from the WARC, and Render runs before
// that). It states the value=English / anchors=original-surface-form contract, the
// required who/what vs optional when/where/how/why split, and the reins submit
// invocation `ccnews submit --key <URL> --in -`. The last failure reason (if any) is
// read from the item's log tail so a retrying agent sees why it failed.
func (ccnewsDef) Render(it *quest.Item) (string, error) {
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
