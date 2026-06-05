//ff:func feature=ingestion type=helper control=selection level=error
//ff:what HEAD로 리소스 메타데이터를 받아 Content-Length를 읽는다. HEAD 미지원·에러·헤더 누락/음수면 ok=false.

package ingest

import (
	"io"
	"net/http"
)

// headSize asks the server for the resource metadata via HEAD and reads
// Content-Length. Returns ok=false when HEAD is unsupported, errors, or the
// header is missing/negative.
func (c *Client) headSize(url string) (int64, bool) {
	req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return 0, false
	}
	req.Header.Set("User-Agent", c.userAgent)
	resp, err := c.http.Do(req)
	if err != nil {
		return 0, false
	}
	defer resp.Body.Close()
	io.Copy(io.Discard, resp.Body)
	if resp.StatusCode != http.StatusOK {
		return 0, false
	}
	if resp.ContentLength >= 0 {
		return resp.ContentLength, true
	}
	return 0, false
}
