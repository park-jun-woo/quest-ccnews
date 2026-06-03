//ff:type feature=article type=model
//ff:what CC-NEWS WARC 파일 내 기사 본문 위치(파일명+오프셋). 본문 자체는 직렬화하지 않는다(WARC가 진실원).

package session

// WARCLoc: location of the article body within a CC-NEWS WARC file. The WARC is
// the source of truth for the body text (body itself is never serialized).
type WARCLoc struct {
	File   string `json:"file"`
	Offset int64  `json:"offset"`
}
