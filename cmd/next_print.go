//ff:func feature=cli type=helper control=sequence
//ff:what next 출력 포매터. 기사 메타·원문 앵커대상 본문·event6 작성 규칙(value=영어, anchors=원문 표면형, 필수 who/when/what)·submit 사용법을 사람/에이전트가 읽을 수 있게 찍는다.

package cmd

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// printNext writes the next-article briefing: identifying metadata, the
// anchor-target body text (the exact text the gate will substring-match against),
// and the event6 authoring contract. The agent reads this, produces an event6
// JSON, and submits it via `ccnews submit`.
func printNext(w io.Writer, a *session.Article, bodyText string) {
	fmt.Fprintln(w, "=== 다음 TODO 기사 ===")
	fmt.Fprintf(w, "URL : %s\n", a.URL)
	fmt.Fprintf(w, "호스트: %s\n", a.Host)
	if a.Lang != "" {
		fmt.Fprintf(w, "언어: %s\n", a.Lang)
	}
	fmt.Fprintf(w, "시도: %d/%d\n", a.Tries, session.MaxTries)
	fmt.Fprintln(w)

	fmt.Fprintln(w, "=== 원문 본문(앵커 대상 텍스트) ===")
	fmt.Fprintln(w, bodyText)
	fmt.Fprintln(w)

	fmt.Fprintln(w, "=== event6 작성 규칙 ===")
	fmt.Fprintln(w, "위 본문에서 육하원칙(event6)을 만들어 제출하세요.")
	fmt.Fprintln(w, "  - value  : 영어로 산출 (날짜는 ISO, 숫자는 정규화)")
	fmt.Fprintln(w, "  - anchors: 원문에 글자 그대로 나타난 표면형 토큰들 (인명·날짜·숫자·지명)")
	fmt.Fprintln(w, "  - 필수: who·when·what (값+앵커 필수, 비우면 FAIL)")
	fmt.Fprintln(w, "  - 선택: where·how·why (있으면 앵커, 앵커 0개면 REVIEW)")
	fmt.Fprintln(w, "  - 게이트는 anchors가 원문 substring인지만 본다. value는 검증 대상 아님.")
	fmt.Fprintln(w, "  - 앵커가 하나라도 원문에 없으면 FAIL(환각). tries 소진 시 DONE.")
	fmt.Fprintln(w)

	fmt.Fprintln(w, "=== 제출 ===")
	fmt.Fprintf(w, "  ccnews submit --url %q --event6 <event6.json>\n", a.URL)
	fmt.Fprintln(w, "  (또는 --event6 - 로 stdin)")
	fmt.Fprintln(w, "  event6 JSON 예: {\"who\":{\"value\":\"...\",\"anchors\":[\"...\"]},"+
		"\"when\":{...},\"what\":{...},\"where\":null,\"how\":null,\"why\":null}")
}
