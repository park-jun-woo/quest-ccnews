//ff:type feature=robots type=model
//ff:what 그룹 내 단일 Allow/Disallow 지시자. * 와일드카드와 끝 $ 앵커를 가질 수 있는 path 패턴(순수 데이터).

package robots

// Rule is a single Allow or Disallow directive within a group.
type Rule struct {
	Allow   bool   // true = Allow, false = Disallow
	Pattern string // raw path pattern, may contain * and trailing $
}
