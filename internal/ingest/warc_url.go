//ff:func feature=ingestion type=helper control=sequence
//ff:what warc.paths 목록의 WARC 객체 경로에 대한 다운로드 URL을 만든다. 순수 함수.

package ingest

import "fmt"

// WarcURL returns the download URL for a WARC object path taken from a
// warc.paths listing. Pure (no IO).
func WarcURL(path string) string {
	return fmt.Sprintf("%s/%s", baseURL, path)
}
