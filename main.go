//ff:func feature=cli type=command control=sequence
//ff:what 프로그램 진입점. cmd 패키지의 루트 명령을 실행한다.

package main

import "github.com/park-jun-woo/quest-ccnews/cmd"

func main() {
	cmd.Execute()
}
