//ff:type feature=robots type=model
//ff:what robots 게이트 패키지 설명. RFC 9309 수준 파서·평가(순수)와 robots.txt 1회 fetch(IO)를 분리한다.

// Package robots implements the ccnews robots gate (Phase004): a one-time
// re-check of each host's current robots.txt and a deterministic per-article
// path evaluation.
//
// Design split (Phase004 §설계원칙): parsing and longest-match evaluation are
// pure functions (Parse, Evaluate, MatchPattern, NormalizePath) so they are
// unit-testable without network; the only IO is Fetch, which retrieves the
// host's robots.txt once.
//
// UA matching (Phase004 §robots규칙): we match against the product token
// "parkjunwoo-quest". No site singles us out, so in practice the "*" group
// applies. We do NOT impersonate CCBot — that respect is deliberately deferred.
//
// Status semantics: a missing robots.txt or any 4xx is treated as allowed
// (status "missing"). A 5xx or transport timeout is "unreachable"; per Phase004
// §열린결정 the conservative policy is encoded in the gate, not the parser.
package robots
