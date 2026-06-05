//ff:func feature=ingestion type=helper control=sequence
//ff:what restoreScratch 단위테스트(에러·재초기화 경로). 잘못된 ingestion·hosts Meta는 디코드 에러를 전파하는지, 디코드된 nil hosts(JSON null)·nil ProcessedWarcs(명시 null)가 빈 map/slice로 재초기화되는지 검증한다.

package runcmd

import (
	"testing"

	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestRestoreScratchErrors(t *testing.T) {
	t.Run("bad ingestion Meta propagates decode error", func(t *testing.T) {
		s := quest.New()
		s.SetMeta(metaIngestion, "not-an-object")
		if _, err := restoreScratch(s, "ua"); err == nil {
			t.Errorf("err = nil, want ingestion decode error")
		}
	})

	t.Run("bad hosts Meta propagates decode error", func(t *testing.T) {
		s := quest.New()
		s.SetMeta(metaHosts, "not-a-map")
		if _, err := restoreScratch(s, "ua"); err == nil {
			t.Errorf("err = nil, want hosts decode error")
		}
	})

	t.Run("decoded nil hosts is reinitialized", func(t *testing.T) {
		s := quest.New()
		// hosts present but JSON null → decodes to nil map, must be reinit'd.
		s.Meta = map[string]any{metaHosts: nil}
		sc, err := restoreScratch(s, "ua")
		if err != nil {
			t.Fatalf("err = %v", err)
		}
		if sc.Hosts == nil {
			t.Errorf("Hosts nil, want reinitialized")
		}
	})

	t.Run("decoded nil ProcessedWarcs is reinitialized", func(t *testing.T) {
		s := quest.New()
		// explicit null overwrites the fresh-scratch empty slice with nil.
		s.SetMeta(metaIngestion, map[string]any{
			"source":          "cc-news",
			"processed_warcs": nil,
		})
		sc, err := restoreScratch(s, "ua")
		if err != nil {
			t.Fatalf("err = %v", err)
		}
		if sc.Ingestion.ProcessedWarcs == nil {
			t.Errorf("ProcessedWarcs nil, want reinitialized")
		}
	})
}
