//ff:type feature=gate type=model
//ff:what 컴파일타임 계약 보증. ccnewsDef가 reins gate.Definition(Seed/Render/Prepare/Rules)을 만족함을 빌드 시점에 강제한다. ccnews는 gate.Evaluator를 구현하지 않으므로 submit은 gate.Evaluate(Rules) 평평한 경로(레벨집계)를 탄다.

package ccnewsquest

import "github.com/park-jun-woo/reins/pkg/gate"

// Compile-time guarantee that ccnewsDef satisfies gate.Definition. ccnews does NOT
// implement gate.Evaluator, so the cli submit wiring takes the flat
// gate.Evaluate(Rules) path (level aggregation), not a defeat graph.
var _ gate.Definition = ccnewsDef{}
