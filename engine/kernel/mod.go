package kernel

const (
	// Author - 囲碁思考エンジンの作者名だぜ☆（＾～＾）
	Author = "Satoshi Takahashi"
)

type KomiType float64

// Komi - コミ。 6.5 といった数字を入れるだけ。実行速度優先で 64bitに。
var Komi KomiType

type MovesNumType int

// MaxMovesNum - 上限手数
var MaxMovesNum MovesNumType
