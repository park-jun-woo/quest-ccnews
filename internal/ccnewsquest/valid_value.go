//ff:func feature=gate type=helper control=sequence
//ff:what 한 필드 value의 내재적 위생을 판정한다(anchor.validValue의 동작보존 이식). 양끝 트림 후 룬 길이가 2 미만이거나 플레이스홀더 블록리스트(대소문자 무시 정확 일치)에 들면 무효(false). 앵커-value 매핑이 아니라 value 자체만 본다. 순수 함수.

package ccnewsquest

import (
	"strings"
	"unicode/utf8"
)

// placeholderValues ports anchor.placeholderValues verbatim: affirmatively-junk
// value tokens a weak model emitted instead of a real fact (Phase009 L3). Matched
// case-insensitively against the trimmed value as an exact string.
var placeholderValues = map[string]struct{}{
	"subject": {}, "event": {}, "unknown": {},
	"n/a": {}, "na": {}, "none": {}, "null": {},
	"tbd": {}, "-": {}, "--": {},
}

// validValue is the behavior-preserving port of anchor.validValue (kept local
// because the original is unexported and Phase012 forbids touching internal/anchor).
// It reports whether a field value is intrinsically usable: INVALID (false) when,
// after trimming surrounding whitespace, it is empty, has a rune length below 2, or
// matches placeholderValues (case-insensitive exact). Value-intrinsic hygiene only —
// it never compares the value to the anchors. Pure.
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
