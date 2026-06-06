package agentcfg

// agentBodyText is the article body the candy WARC carries; the stub anchors come
// from it (e.g. "Reporter") so the real anchor gate PASSes on a faithful event6.
const agentBodyText = "This is a sufficiently long article body that exceeds the minimum length " +
	"threshold required by the trust gate so that the article passes the body-length " +
	"check and is not skipped as too short for anchoring purposes. Reporter wrote it."

// agentPassHTML clears the trust gate (structured JSON-LD + self-declared title +
// body >= MinBodyLen) so Prepare reaches the anchor gate.
const agentPassHTML = `<html><head>` +
	`<script type="application/ld+json">{"@type":"NewsArticle","headline":"Framework Agreement","author":{"name":"Reporter"},"datePublished":"2026-06-05","articleBody":"` +
	agentBodyText +
	`","inLanguage":"en"}</script>` +
	`</head><body></body></html>`

// goodEvent6 is a faithful event6 whose anchors are verbatim substrings of
// agentBodyText, so the anchor gate PASSes.
const goodEvent6 = `{"who":{"value":"Reporter","anchors":["Reporter"]},` +
	`"what":{"value":"wrote an article","anchors":["wrote"]}}`

// fencedGoodEvent6 wraps goodEvent6 in a markdown fence to exercise the lenient
// decode rescue (Phase015 A).
const fencedGoodEvent6 = "```json\n" + goodEvent6 + "\n```"

// hallucEvent6 anchors a token absent from agentBodyText, so required-anchor-real FAILs.
const hallucEvent6 = `{"who":{"value":"Ghost","anchors":["Nonexistent-Anchor-XYZ"]},` +
	`"what":{"value":"wrote an article","anchors":["wrote"]}}`
