//ff:func feature=cli type=command control=sequence level=error
//ff:what 프로그램 진입점. reins cli.NewQuestCmd로 ccnews 퀘스트 CLI(scan/next/submit/status/export/rules)를 조립하고, 투트랙 WARC 인제스천 `run` 명령을 ExtraCommand로 부착해 실행한다. 도메인 로직은 ccnewsquest.Def(userAgent, cacheDir) 하나, 인제스천은 runcmd.New() 하나.

package main

import (
	"os"

	"github.com/park-jun-woo/quest-ccnews/internal/ccnewsquest"
	"github.com/park-jun-woo/quest-ccnews/internal/runcmd"
	"github.com/park-jun-woo/reins/pkg/cli"
	"github.com/spf13/cobra"
)

// defaultUserAgent matches Phase001 결정 4 (robots UA).
const defaultUserAgent = "parkjunwoo-quest/0.1 (+https://www.parkjunwoo.com)"

// defaultCacheDir is where downloaded .warc.gz files are cached on disk; it is the
// same directory Prepare's WARC re-read client reads from, keeping run and submit
// consistent (Phase013).
const defaultCacheDir = "warc-cache"

func main() {
	root := cli.NewQuestCmd("ccnews", ccnewsquest.Def(defaultUserAgent, defaultCacheDir), cli.Options{
		Version:       "0.3",
		ExtraCommands: []*cobra.Command{runcmd.New(defaultUserAgent, defaultCacheDir)},
	})
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
