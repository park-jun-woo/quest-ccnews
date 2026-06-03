//ff:func feature=ingestion type=helper control=sequence
//ff:what WARC basename(CC-NEWS-YYYYMMDDhhmmss-...)에서 연/월을 파싱한다. 형식이 안 맞으면 ok=false. 순수 함수.

package ingest

import "strconv"

// MonthFromWarc parses the dump month from a CC-NEWS WARC basename of the form
// "CC-NEWS-YYYYMMDDhhmmss-NNNNN.warc.gz". Returns ok=false when the name does not
// carry a parseable YYYYMM timestamp. Pure (no IO).
func MonthFromWarc(name string) (Month, bool) {
	const prefix = "CC-NEWS-"
	if len(name) < len(prefix)+6 || name[:len(prefix)] != prefix {
		return Month{}, false
	}
	ts := name[len(prefix):]
	year, err := strconv.Atoi(ts[0:4])
	if err != nil {
		return Month{}, false
	}
	month, err := strconv.Atoi(ts[4:6])
	if err != nil || month < 1 || month > 12 {
		return Month{}, false
	}
	return Month{Year: year, Month: month}, true
}
