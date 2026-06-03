//ff:func feature=output type=helper control=iteration dimension=1 level=error
//ff:what 세션의 종단(PASS/REVIEW/BLOCKED/SKIPPED) 미emit 기사를 훑어 out 파일에 한 줄씩 append하고 Emitted=true로 표시한다. emit-once 보장. 추가된 건수를 반환.

package output

import "github.com/park-jun-woo/quest-ccnews/internal/session"

// Sweep appends a JSONL record for every terminal-state article in s that has
// not yet been emitted, marking each Emitted=true after a successful append so
// the same article is never written twice (emit-once). It mutates s.Articles
// (the Emitted flags) but does not persist the session — the caller saves. The
// host cache entry is looked up by Article.Host (may be absent). Returns the
// number of records appended. On the first append error it returns that error
// with the count written so far; already-flagged articles stay flagged.
func Sweep(s *session.Session, path string) (int, error) {
	if s == nil {
		return 0, nil
	}
	written := 0
	for _, a := range s.Articles {
		if a.Emitted {
			continue
		}
		rec := Render(a, s.Hosts[a.Host])
		if rec == nil {
			continue // not a terminal state
		}
		if err := Append(path, rec); err != nil {
			return written, err
		}
		a.Emitted = true
		written++
	}
	return written, nil
}
