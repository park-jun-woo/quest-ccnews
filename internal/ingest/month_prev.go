//ff:func feature=ingestion type=helper control=sequence
//ff:what 한 달 이전 Month를 반환하고, CC-NEWS 최古(2016-08)에 도달했는지 알린다. 순수 함수.

package ingest

// Prev returns the month immediately before m, and ok=false if m is already the
// earliest CC-NEWS dump (2016-08) — i.e. there is no earlier month. Pure (no IO).
func (m Month) Prev() (Month, bool) {
	if m.Year == firstYear && m.Month == firstMonth {
		return m, false
	}
	if m.Month == 1 {
		return Month{Year: m.Year - 1, Month: 12}, true
	}
	return Month{Year: m.Year, Month: m.Month - 1}, true
}
