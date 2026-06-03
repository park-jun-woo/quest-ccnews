//ff:type feature=robots type=model
//ff:what robots.txt fetch 클라이언트. HTTP 클라이언트와 UA 헤더 문자열, UA 매칭용 product token을 묶은 얇은 IO 핸들.

package robots

import "net/http"

// Client is the thin IO handle for the one-time robots.txt re-check: an HTTP
// client, the full User-Agent header string to send, and the product token used
// for group matching (the leading token of the UA, e.g. "parkjunwoo-quest").
type Client struct {
	http         *http.Client
	userAgent    string
	productToken string
}
