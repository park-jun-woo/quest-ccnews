//ff:func feature=ingestion type=helper control=sequence
//ff:what 한 Month에 대한 CC-NEWS warc.paths.gz의 다운로드 URL을 만든다. 순수 함수.

package ingest

import "fmt"

// baseURL is the Common Crawl public data host for CC-NEWS.
const baseURL = "https://data.commoncrawl.org"

// PathsURL returns the warc.paths.gz URL for a CC-NEWS monthly dump, e.g.
// "https://data.commoncrawl.org/crawl-data/CC-NEWS/2026/06/warc.paths.gz".
// Pure (no IO).
func PathsURL(m Month) string {
	return fmt.Sprintf("%s/crawl-data/CC-NEWS/%s/warc.paths.gz", baseURL, m)
}
