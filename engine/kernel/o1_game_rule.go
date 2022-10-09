package kernel

type KomiFloat float64

// Komi - コミ。 6.5 といった数字を入れるだけ。実行速度優先で 64bitに。
var Komi KomiFloat

type PositionNumberInt int

// MaxPositionNumber - 上限手数
var MaxPositionNumber PositionNumberInt
