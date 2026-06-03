//ff:func feature=anchor type=helper control=sequence
//ff:what 한 필드 value의 내재적 위생을 판정한다. 양끝 트림 후 비었거나, 룬 길이가 2 미만이거나, 플레이스홀더 블록리스트(대소문자 무시 정확 일치)에 들면 무효(false). 앵커-value 매핑이 아니라 value 자체만 본다. 순수 함수.

package anchor

import (
	"strings"
	"unicode/utf8"
)

// placeholderValues are affirmatively-junk value tokens a weak model emitted
// instead of a real fact (Phase009 L3, 벡터 5: who="Subject", value="Event").
// Matched case-insensitively against the trimmed value as an exact string —
// conservative, so no real word is over-blocked. The plan's stated set
// (Subject/Event/Unknown/N/A/-) is extended with obvious equivalents.
var placeholderValues = map[string]struct{}{
	"subject": {}, "event": {}, "unknown": {},
	"n/a": {}, "na": {}, "none": {}, "null": {},
	"tbd": {}, "-": {}, "--": {},
}

// validValue reports whether a field value is intrinsically usable. It is INVALID
// (returns false) when, after trimming surrounding whitespace, it is empty, has a
// rune length below 2, or matches placeholderValues (case-insensitive exact). This
// is value-intrinsic hygiene only — it never compares the value to the anchors, so
// the package's no anchor→value mapping principle holds. Pure.
func validValue(v string) bool {
	t := strings.TrimSpace(v)
	if utf8.RuneCountInString(t) < 2 {
		return false
	}
	if _, ok := placeholderValues[strings.ToLower(t)]; ok {
		return false
	}
	return true
}
