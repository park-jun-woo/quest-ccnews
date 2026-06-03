package extract

// longBody is a body string comfortably above MinBodyLen (200 bytes). Shared
// test fixture for gate/parse/apply tests. Const-only file (F5).
const longBody = "This is a sufficiently long article body that exceeds the minimum length " +
	"threshold required by the trust gate so that the article passes the body-length " +
	"check and is not skipped as too short for anchoring purposes during phase six. " +
	"It contains several sentences of running prose."
