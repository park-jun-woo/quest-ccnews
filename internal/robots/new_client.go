//ff:func feature=robots type=helper control=sequence
//ff:what 기본 타임아웃을 가진 robots fetch 클라이언트를 만든다. UA 헤더에서 product token(첫 토큰)을 뽑아 매칭에 쓴다.

package robots

import (
	"net/http"
	"time"
)

// NewClient builds a robots client with a short default timeout (robots.txt is
// small and a slow host should fall to "unreachable", not stall the run). The
// product token used for UA-group matching is the leading token of userAgent
// before any "/" or space — e.g. "parkjunwoo-quest/0.1 (...)" → "parkjunwoo-quest".
func NewClient(userAgent string) *Client {
	return &Client{
		http:         &http.Client{Timeout: 15 * time.Second},
		userAgent:    userAgent,
		productToken: ProductToken(userAgent),
	}
}
