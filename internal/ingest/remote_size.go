//ff:func feature=ingestion type=helper control=selection level=error
//ff:what WARC의 바이트 크기를 본문 다운로드 없이 탐침한다. HEAD를 먼저 시도하고, 실패하면 단일 바이트 Range GET으로 폴백한다. 비음수 크기를 정한 경우에만 (size,true).

package ingest

// remoteSize probes the WARC's byte size without downloading the body. It tries
// a HEAD request first; if HEAD is unsupported or yields no usable
// Content-Length, it falls back to a Range GET asking for a single byte and
// reading the size off Content-Range. It returns (size, true) only when a
// non-negative size is determined. The probe never returns a body to copy, so a
// cache hit costs at most a few bytes over the wire.
func (c *Client) remoteSize(objectPath string) (int64, bool) {
	url := WarcURL(objectPath)

	if n, ok := c.headSize(url); ok {
		return n, true
	}
	return c.rangeSize(url)
}
