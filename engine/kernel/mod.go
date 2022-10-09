package kernel

const (
	// Author - 囲碁思考エンジンの作者名だぜ☆（＾～＾）
	Author = "Satoshi Takahashi"
)

func SetBoardSize(boardSize int) {
	BoardSize = boardSize
	BoardArea = BoardSize * BoardSize
}

// BoardSize - 何路盤
var BoardSize int

// BoardArea - 壁無し盤の面積
var BoardArea int

// GetMemoryBoardWidth - 枠付きの盤の一辺の交点数
func GetMemoryBoardWidth() int {
	return BoardSize + 2
}

// GetMemoryBoardArea - 壁付き盤の面積
func GetMemoryBoardArea() int {
	return GetMemoryBoardWidth() * GetMemoryBoardWidth()
}

type KomiType float64

// Komi - コミ。 6.5 といった数字を入れるだけ。実行速度優先で 64bitに。
var Komi KomiType

type MovesNumType int

// MaxMovesNum - 上限手数
var MaxMovesNum MovesNumType

// Point - 交点の座標。壁を含む盤の左上を 0 とします
type Point int

// Pass - パス
const Pass Point = 0

// Dir4 - ４方向（東、北、西、南）の番地。初期値は仮の値。 2015年講習会サンプル、GoGo とは順序が違います
var Dir4 = [4]Point{1, -9, -1, 9}

type Direction4 int

// Dir4に対応
const (
	East Direction4 = iota
	North
	West
	South
)
