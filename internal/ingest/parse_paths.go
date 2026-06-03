//ff:func feature=ingestion type=helper control=iteration dimension=1
//ff:what CC-NEWS warc.paths 본문(개행 구분 경로 목록)을 파싱해 비어있지 않은 경로 슬라이스로 만든다. 순수 함수.

package ingest

import (
	"bufio"
	"strings"
)

// ParsePaths parses the decompressed body of a CC-NEWS warc.paths file into a
// slice of WARC object paths (e.g. "crawl-data/CC-NEWS/2026/06/CC-NEWS-...warc.gz").
// Blank lines and surrounding whitespace are stripped. Pure (no IO).
func ParsePaths(body string) []string {
	var paths []string
	sc := bufio.NewScanner(strings.NewReader(body))
	sc.Buffer(make([]byte, 64*1024), 1024*1024)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		paths = append(paths, line)
	}
	return paths
}
