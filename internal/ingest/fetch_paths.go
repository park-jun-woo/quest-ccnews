//ff:func feature=ingestion type=helper control=sequence level=error
//ff:what 한 Month의 warc.paths.gz를 받아 gunzip 후 경로 목록으로 파싱한다(얇은 IO + 순수 ParsePaths 호출).

package ingest

import (
	"compress/gzip"
	"io"
)

// FetchPaths downloads the CC-NEWS warc.paths.gz for a month, gunzips it, and
// returns the WARC object paths. IO is kept thin: the gunzipped body is handed to
// the pure ParsePaths. Returns an empty slice (no error) is impossible — a
// missing month surfaces as an httpGet error.
func (c *Client) FetchPaths(m Month) ([]string, error) {
	body, err := httpGet(c.http, PathsURL(m), c.userAgent)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	gz, err := gzip.NewReader(body)
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	raw, err := io.ReadAll(gz)
	if err != nil {
		return nil, err
	}
	return ParsePaths(string(raw)), nil
}
