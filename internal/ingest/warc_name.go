//ff:func feature=ingestion type=helper control=sequence
//ff:what CC-NEWS WARC 객체 경로에서 잠금/커서 키로 쓰는 basename(예: CC-NEWS-...warc.gz)을 추출한다. 순수 함수.

package ingest

import "strings"

// WarcName returns the basename used as the cursor / processed_warcs lock key
// for a CC-NEWS WARC object path. For "crawl-data/CC-NEWS/2026/06/CC-NEWS-x.warc.gz"
// it returns "CC-NEWS-x.warc.gz". Pure (no IO).
func WarcName(path string) string {
	if i := strings.LastIndex(path, "/"); i >= 0 {
		return path[i+1:]
	}
	return path
}
