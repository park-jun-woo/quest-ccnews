//ff:type feature=extract type=model
//ff:what extract 패키지 개요 — WARC HTML 바이트에서 구조화 데이터(JSON-LD/OG)만으로 기사 필드를 결정론적 추출하고 신뢰 게이트를 판정한다.

// Package extract turns one article's raw WARC HTML bytes into a structured
// extraction result plus a deterministic trust decision. It is a pure-parse
// layer: HTML bytes in, result out, no network and no AI.
//
// Extraction trusts only self-declared structured data, in priority order:
//
//  1. JSON-LD  <script type="application/ld+json"> with @type NewsArticle/Article:
//     headline → title, articleBody → body, author(.name|string) → author,
//     datePublished → published_at, publisher.name → media_name, inLanguage → lang.
//  2. OG/meta fallback (only when JSON-LD yields no Article): og:title,
//     og:site_name → media_name, article:published_time → published_at,
//     article:author → author.
//
// The anchor-target body text (Phase006 input) is articleBody when present,
// else the tag-stripped full HTML text. That text is returned to the caller but
// never stored in the session (the WARC stays the source); only its length is
// kept as Extracted.BodyLen.
//
// Trust gate (deterministic): PASS requires (1) a non-empty title from a
// structured source and (2) anchor-target text length >= MinBodyLen. Otherwise
// SKIPPED, with a reason distinguishing "no structured data" from "body too
// short for anchoring".
package extract
