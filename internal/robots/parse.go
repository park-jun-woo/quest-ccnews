//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what robots.txt 바이트를 RFC 9309 룰셋으로 파싱한다. User-agent 그룹·Allow/Disallow·crawl-delay를 모은다. 순수 함수(네트워크 0).

package robots

import (
	"bufio"
	"bytes"
)

// Parse turns robots.txt content into a Ruleset following RFC 9309 grouping:
// consecutive user-agent lines start (or extend) a group, and the rules that
// follow apply to all those agents until the next user-agent line begins a new
// group. Comments (#...) and unknown fields are ignored. Pure — no IO.
func Parse(content []byte) *Ruleset {
	p := &parser{rs: &Ruleset{}}
	sc := bufio.NewScanner(bytes.NewReader(content))
	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for sc.Scan() {
		field, value, ok := parseLine(sc.Text())
		if ok {
			p.feed(field, value)
		}
	}
	return p.rs
}
