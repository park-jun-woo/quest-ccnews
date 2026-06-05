//ff:func feature=ingestion type=helper control=selection level=error
//ff:what HEAD 미지원 시 단일 바이트 Range GET으로 폴백해 Content-Range("bytes 0-0/<total>")에서 전체 크기를 읽는다. 에러·파싱불가면 ok=false.

package ingest

import (
	"fmt"
	"io"
	"net/http"
)

// rangeSize falls back to a single-byte Range GET and reads the total size off
// the Content-Range header ("bytes 0-0/<total>"). Used when HEAD is
// unsupported. Returns ok=false on any error or unparsable response.
func (c *Client) rangeSize(url string) (int64, bool) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, false
	}
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Range", "bytes=0-0")
	resp, err := c.http.Do(req)
	if err != nil {
		return 0, false
	}
	defer resp.Body.Close()
	io.Copy(io.Discard, resp.Body)
	if resp.StatusCode != http.StatusPartialContent {
		return 0, false
	}
	cr := resp.Header.Get("Content-Range")
	var total int64
	// Expected form: "bytes 0-0/123"
	if _, err := fmt.Sscanf(cr, "bytes 0-0/%d", &total); err != nil {
		return 0, false
	}
	if total < 0 {
		return 0, false
	}
	return total, true
}
