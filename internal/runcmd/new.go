//ff:type feature=ingestion type=model
//ff:func feature=ingestion type=command control=sequence
//ff:what options는 run 명령의 해석된 설정·플래그 타깃을 묶는다. New는 reins cli에 부착할 `run` ExtraCommand를 만든다(G1). 투트랙 WARC 인제스천 루프를 reins 세션 위에서 돈다. userAgent·cacheDir는 ccnewsquest.Def와 같은 값을 받아 run/세션 Meta와 일관되게 한다. --track/--max-warcs/--cache-dir 플래그를 달고, 세션 경로는 root의 persistent --session을 상속해 읽는다.

package runcmd

import "github.com/spf13/cobra"

// options holds the run command's resolved configuration and flag targets.
type options struct {
	userAgent string
	cacheDir  string
	track     string
	maxWarcs  int
	robots    bool           // when false, skip the eager per-host robots.txt fetch in bridge
	cmd       *cobra.Command // set in New so sessionPath() can read the inherited flag
}

// sessionPath returns the session file path from the root's persistent --session
// flag (inherited by the run subcommand), defaulting to "session.json".
func (o *options) sessionPath() string {
	if o.cmd != nil {
		if f := o.cmd.Flags().Lookup("session"); f != nil {
			return f.Value.String()
		}
	}
	return "session.json"
}

// New builds the `run` ExtraCommand for reins cli.Options.ExtraCommands. userAgent
// and cacheDir mirror ccnewsquest.Def so the WARC cache and crawl UA stay
// consistent between ingestion (run) and re-read (submit's Prepare).
func New(userAgent, cacheDir string) *cobra.Command {
	o := &options{userAgent: userAgent}
	cmd := &cobra.Command{
		Use:   "run",
		Short: "투트랙 WARC 인제스천 루프(다운로드→레코드→reins Item 시드→커서 전진)",
		Long: `CC-NEWS WARC를 받아 response 레코드를 TODO 기사 Item으로 reins 세션에 시드한다.

  --track forward   최신 덤프만 따라잡고 따라잡으면 waiting
  --track backward  과거로 내려가며 2016-08(exhausted)까지
  --track both      둘 다(기본)

커서·processed·호스트 robots 캐시는 세션 Meta에 보존되어 다음 run에 이어진다.
robots 거부 기사는 BLOCKED로 직접 시드한다.`,
		RunE: o.run,
	}
	o.cmd = cmd
	cmd.Flags().StringVar(&o.track, "track", "both", "처리할 트랙: forward|backward|both")
	cmd.Flags().IntVar(&o.maxWarcs, "max-warcs", 0, "처리할 WARC 최대 개수(0=무제한)")
	cmd.Flags().StringVar(&o.cacheDir, "cache-dir", cacheDir, "다운로드 WARC 캐시 디렉터리")
	cmd.Flags().BoolVar(&o.robots, "robots", true, "브리지 시 호스트별 robots.txt를 1회 fetch해 거부 기사를 BLOCKED로 시드(false면 생략하고 모두 TODO 시드 — 대량 인제스천 시 1892개 호스트 라이브 fetch 폭주 회피)")
	return cmd
}
