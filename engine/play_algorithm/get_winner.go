package play_algorithm

import (
	e "github.com/muzudho/kifuwarabe-go-toramedal/engine/kernel"
)

// WrapGettingOfWinner - 盤を束縛変数として与えます
func WrapGettingOfWinner(position *e.Position) *func(turnColor e.Stone) int {
	// 「手番の勝ちなら1、引き分けなら0、手番の負けなら-1を返す関数（自分視点）」を作成します
	// * `turnColor` - 手番の石の色
	var getWinner = func(turnColor e.Stone) int {
		return getWinner(position, turnColor)
	}

	return &getWinner
}

// 手番の勝ちなら1、引き分けなら0、手番の負けなら-1（自分視点）
// * `turnColor` - 手番の石の色
func getWinner(position *e.Position, turnColor e.Stone) int {
	var mk = [4]int{}
	var kind = [3]int{0, 0, 0}
	var score, blackArea, whiteArea, blackSum, whiteSum int

	var setPoint = func(z e.Point) {
		var color2 = position.GetBoard().GetStoneAt(z)
		kind[color2]++
		if color2 == 0 {
			mk[1] = 0
			mk[2] = 0
			for dir := e.Cell_4Directions(0); dir < 4; dir++ {
				mk[position.GetBoard().GetStoneAt(z+position.GetBoard().GetCoordinate().GetRelativePointOf(dir))]++
			}
			if mk[1] != 0 && mk[2] == 0 {
				blackArea++
			}
			if mk[2] != 0 && mk[1] == 0 {
				whiteArea++
			}
		}
	}

	position.GetBoard().GetCoordinate().ForeachCellWithoutWall(setPoint)

	blackSum = kind[1] + blackArea
	whiteSum = kind[2] + whiteArea
	score = blackSum - whiteSum
	var win = 0
	if 0 < e.KomiFloat(score)-position.GetBoard().GetGameRule().GetKomi() {
		win = 1
	}
	if turnColor == 2 {
		win = -win
	} // gogo07

	return win
}
