//ff:func feature=ingestion type=helper control=sequence
//ff:what time.Time를 CC-NEWS Month(연/월)로 변환한다(UTC 기준). 순수 함수.

package ingest

import "time"

// MonthOf returns the CC-NEWS Month for an instant, in UTC. Pure (no IO).
func MonthOf(t time.Time) Month {
	u := t.UTC()
	return Month{Year: u.Year(), Month: int(u.Month())}
}
