package kernel

// KomiFloat - コミ。 6.5 といった数字を入れるだけ。実行速度優先で 64bitに
type KomiFloat float64

// PositionNumberInt - 何手目
type PositionNumberInt int

// GameRule - 対局ルール
type GameRule struct {
	// コミ。 6.5 といった数字を入れるだけ。実行速度優先で 64bitに
	komi KomiFloat

	// 上限手数
	maxPositionNumber PositionNumberInt
}

// NewGameRule - 新規作成
func NewGameRule(komi KomiFloat, maxPositionNumber PositionNumberInt) *GameRule {
	var gr = new(GameRule)

	gr.komi = komi
	gr.maxPositionNumber = maxPositionNumber

	return gr
}

// GetKomi - コミ取得
func (gr *GameRule) GetKomi() KomiFloat {
	return gr.komi
}

// SetKomi - コミ設定
func (gr *GameRule) SetKomi(komi KomiFloat) {
	gr.komi = komi
}

// GetMaxPositionNumber - 上限手数
func (gr *GameRule) GetMaxPositionNumber() PositionNumberInt {
	return gr.maxPositionNumber
}
