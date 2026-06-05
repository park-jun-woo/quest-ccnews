//ff:func feature=ingestion type=helper control=sequence
//ff:what restoreScratch 단위테스트(정상 경로). 빈 Meta면 defaultUA·초기화된 Hosts/ProcessedWarcs의 빈 스크래치, 저장 UA가 기본을 덮어쓰되 비문자열·빈 문자열은 기본 유지, ingestion·hosts가 Meta에서 디코드되는지 검증한다. 에러·재초기화 경로는 TestRestoreScratchErrors에서 검증한다.

package runcmd

import (
	"testing"

	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestRestoreScratch(t *testing.T) {
	t.Run("empty Meta yields fresh scratch with defaultUA", func(t *testing.T) {
		s := quest.New()
		sc, err := restoreScratch(s, "default-ua")
		if err != nil {
			t.Fatalf("err = %v", err)
		}
		if sc.UserAgent != "default-ua" {
			t.Errorf("UserAgent = %q, want default-ua", sc.UserAgent)
		}
		if sc.Hosts == nil {
			t.Errorf("Hosts nil, want initialized map")
		}
		if sc.Ingestion.ProcessedWarcs == nil {
			t.Errorf("ProcessedWarcs nil, want initialized slice")
		}
		if len(sc.Articles) != 0 {
			t.Errorf("Articles = %d, want 0", len(sc.Articles))
		}
	})

	t.Run("stored UA overrides default", func(t *testing.T) {
		s := quest.New()
		s.SetMeta(metaUserAgent, "stored-ua")
		sc, err := restoreScratch(s, "default-ua")
		if err != nil {
			t.Fatalf("err = %v", err)
		}
		if sc.UserAgent != "stored-ua" {
			t.Errorf("UserAgent = %q, want stored-ua", sc.UserAgent)
		}
	})

	t.Run("non-string or empty UA keeps default", func(t *testing.T) {
		for name, v := range map[string]any{"empty": "", "wrong-type": 42} {
			t.Run(name, func(t *testing.T) {
				s := quest.New()
				s.SetMeta(metaUserAgent, v)
				sc, err := restoreScratch(s, "default-ua")
				if err != nil {
					t.Fatalf("err = %v", err)
				}
				if sc.UserAgent != "default-ua" {
					t.Errorf("UserAgent = %q, want default-ua", sc.UserAgent)
				}
			})
		}
	})

	t.Run("ingestion and hosts decode from Meta", func(t *testing.T) {
		s := quest.New()
		s.SetMeta(metaIngestion, map[string]any{
			"source":          "cc-news",
			"processed_warcs": []any{"w1"},
		})
		s.SetMeta(metaHosts, map[string]any{
			"ex.com": map[string]any{"media_name": "Ex"},
		})
		sc, err := restoreScratch(s, "ua")
		if err != nil {
			t.Fatalf("err = %v", err)
		}
		if sc.Ingestion.Source != "cc-news" || len(sc.Ingestion.ProcessedWarcs) != 1 {
			t.Errorf("Ingestion = %+v", sc.Ingestion)
		}
		if h := sc.Hosts["ex.com"]; h == nil || h.MediaName != "Ex" {
			t.Errorf("Hosts[ex.com] = %+v", h)
		}
	})
}
