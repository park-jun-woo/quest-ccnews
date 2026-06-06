//ff:func feature=cli type=command control=sequence level=error
//ff:what 프로그램 진입점. reins cli.NewQuestCmd로 ccnews 퀘스트 CLI(scan/next/submit/status/export/rules)를 조립하고, 투트랙 WARC 인제스천 `run` 명령을 ExtraCommand로 부착해 실행한다. 도메인 로직은 ccnewsquest.Def(UserAgent, CacheDir) 하나, 인제스천은 runcmd.New() 하나, 에이전트 설정은 agentcfg 하나.

package main

import (
	"os"

	"github.com/park-jun-woo/quest-ccnews/internal/agentcfg"
	"github.com/park-jun-woo/quest-ccnews/internal/ccnewsquest"
	"github.com/park-jun-woo/quest-ccnews/internal/runcmd"
	"github.com/park-jun-woo/reins/pkg/cli"
	"github.com/spf13/cobra"
)

func main() {
	root := cli.NewQuestCmd("ccnews", ccnewsquest.Def(agentcfg.UserAgent, agentcfg.CacheDir), cli.Options{
		Version:       "0.3",
		ExtraCommands: []*cobra.Command{runcmd.New(agentcfg.UserAgent, agentcfg.CacheDir)},
		Agent: &cli.AgentOptions{
			System:     agentcfg.System,
			RuleSystem: agentcfg.RuleCoaching,
		},
	})
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
