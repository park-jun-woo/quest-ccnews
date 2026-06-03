//ff:type feature=article type=model
//ff:what 구조화 데이터 추출 결과. 본문 텍스트는 저장하지 않고 body_len만 보관한다(재처리는 WARC에서).

package session

// Extracted: structured-data extraction result. The body text is not stored;
// only body_len is kept (the WARC remains the source for re-processing).
type Extracted struct {
	Title       string `json:"title,omitempty"`
	Author      string `json:"author,omitempty"`
	PublishedAt string `json:"published_at,omitempty"`
	Source      string `json:"source,omitempty"` // jsonld|og
	BodyLen     int    `json:"body_len"`
}
