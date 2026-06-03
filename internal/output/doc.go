//ff:type feature=output type=model
//ff:what output 패키지 개요 — 종단(PASS/REVIEW/BLOCKED/SKIPPED) 기사를 JSONL 레코드로 렌더링(순수)하고 emit-once로 out 파일에 한 줄씩 append(IO)한다.

// Package output renders terminal-state articles into JSONL records and appends
// them, one line per article, to the --out file (Phase007). Record rendering and
// the anchor_summary computation are pure; only the file append is IO.
package output
