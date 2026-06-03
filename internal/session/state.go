//ff:type feature=session type=model
//ff:what 기사 퀘스트 상태값 타입과 6개 상태 상수(TODO/PASS/REVIEW/DONE/BLOCKED/SKIPPED).

// Package session manages the runtime state SSOT (session.json) for ccnews:
// the two-track ingestion cursor, the host cache (robots + media info), and the
// article quest work-list with its single-direction state machine.
//
// Article state machine (single direction, irreversible once locked):
//
//	TODO ─┬─► PASS      required (who/when/what) anchors complete
//	      ├─► REVIEW    interpretive fields (how/why) need human review
//	      ├─► DONE       MaxTries exhausted (extraction/anchor failed)
//	      ├─► BLOCKED   robots refused (skip_reason recorded)
//	      └─► SKIPPED   no structured data — untrusted (skip_reason recorded)
//
// PASS/REVIEW/DONE/BLOCKED/SKIPPED are locked (irreversible). NextTODO only
// picks TODO articles.
package session

// State is an article quest state.
type State string

const (
	TODO    State = "TODO"    // not yet processed
	PASS    State = "PASS"    // required anchors complete → locked, irreversible
	REVIEW  State = "REVIEW"  // interpretive fields need human review → locked
	DONE    State = "DONE"    // tries exhausted (extraction/anchor failed) → locked
	BLOCKED State = "BLOCKED" // robots refused → locked (skip_reason recorded)
	SKIPPED State = "SKIPPED" // no structured data, untrusted → locked (skip_reason recorded)
)
