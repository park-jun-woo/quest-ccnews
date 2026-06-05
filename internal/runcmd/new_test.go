//ff:func feature=ingestion type=command control=sequence
//ff:what New 단위테스트. `run` ExtraCommand의 Use="run", RunE 배선, 그리고 --track(기본 both)·--max-warcs(기본 0)·--cache-dir(전달 기본값) 플래그가 올바로 달렸는지 검증한다.

package runcmd

import "testing"

func TestNew(t *testing.T) {
	cmd := New("crawl-ua/0.1", "/tmp/cache")

	if cmd.Use != "run" {
		t.Errorf("Use = %q, want run", cmd.Use)
	}
	if cmd.RunE == nil {
		t.Errorf("RunE not wired")
	}

	flags := cmd.Flags()
	if f := flags.Lookup("track"); f == nil || f.DefValue != "both" {
		t.Errorf("track flag = %+v, want default both", f)
	}
	if f := flags.Lookup("max-warcs"); f == nil || f.DefValue != "0" {
		t.Errorf("max-warcs flag = %+v, want default 0", f)
	}
	if f := flags.Lookup("cache-dir"); f == nil || f.DefValue != "/tmp/cache" {
		t.Errorf("cache-dir flag = %+v, want default /tmp/cache", f)
	}
}
