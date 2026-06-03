//ff:type feature=article type=model
//ff:what 한 번의 submit 시도 기록. Reason은 의견이 아니라 사실 그대로.

package session

// Attempt: one submit attempt record. Reason is the fact, not an opinion.
type Attempt struct {
	Try     int    `json:"try"`
	Verdict string `json:"verdict"`
	Reason  string `json:"reason"`
}
