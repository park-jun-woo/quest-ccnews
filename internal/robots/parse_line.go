//ff:func feature=robots type=helper control=sequence
//ff:what robots.txt 한 줄을 소문자 field와 raw value로 분리한다. 주석 제거, 빈 줄/주석/형식오류는 ok=false. 순수 함수.

package robots

import "strings"

// parseLine splits one robots.txt line into a lowercased field name and its raw
// value, stripping comments. ok=false for blank/comment/malformed lines.
func parseLine(line string) (field, value string, ok bool) {
	if i := strings.IndexByte(line, '#'); i >= 0 {
		line = line[:i]
	}
	line = strings.TrimSpace(line)
	if line == "" {
		return "", "", false
	}
	colon := strings.IndexByte(line, ':')
	if colon < 0 {
		return "", "", false
	}
	field = strings.ToLower(strings.TrimSpace(line[:colon]))
	value = strings.TrimSpace(line[colon+1:])
	if field == "" {
		return "", "", false
	}
	return field, value, true
}
