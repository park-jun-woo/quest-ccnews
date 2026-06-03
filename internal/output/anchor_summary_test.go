//ff:func feature=output type=helper control=iteration dimension=1
//ff:what anchorSummary가 nil/빈/전체앵커/일부앵커/무앵커 케이스에서 "anchored/present" 카운트를 올바르게 내는지 검증한다.

package output

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestAnchorSummary(t *testing.T) {
	tests := []struct {
		name string
		ev   *session.Event6
		want string
	}{
		{"nil", nil, "0/0"},
		{"empty", &session.Event6{}, "0/0"},
		{
			"all anchored",
			&session.Event6{
				Who:  &session.Field{Anchors: []string{"a"}},
				When: &session.Field{Anchors: []string{"b"}},
			},
			"2/2",
		},
		{
			"some anchored",
			&session.Event6{
				Who:   &session.Field{Anchors: []string{"a"}},
				When:  &session.Field{}, // present, no anchor
				Where: &session.Field{Anchors: []string{"c"}},
			},
			"2/3",
		},
		{
			"none anchored",
			&session.Event6{
				What: &session.Field{},
				How:  &session.Field{},
				Why:  &session.Field{},
			},
			"0/3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := anchorSummary(tt.ev); got != tt.want {
				t.Errorf("anchorSummary() = %q, want %q", got, tt.want)
			}
		})
	}
}
