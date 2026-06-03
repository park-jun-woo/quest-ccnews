//ff:type feature=ingestion type=model
//ff:what CC-NEWS 월별 덤프 식별자(연/월). warc.paths URL 구성과 트랙 진행의 단위.

package ingest

import "fmt"

// firstYear / firstMonth: CC-NEWS earliest dump is 2016-08.
const (
	firstYear  = 2016
	firstMonth = 8
)

// Month identifies one CC-NEWS monthly dump. It is the unit of warc.paths URL
// construction and of track progression.
type Month struct {
	Year  int
	Month int // 1..12
}

// String formats the month as "YYYY/MM" (the path segment used by CC-NEWS).
func (m Month) String() string {
	return fmt.Sprintf("%04d/%02d", m.Year, m.Month)
}
