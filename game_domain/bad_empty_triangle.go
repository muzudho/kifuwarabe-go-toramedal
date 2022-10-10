package game_domain

// Empty triangle (アキ三角)
//
// x.
// xx

import (
	e "github.com/muzudho/kifuwarabe-go-toramedal/engine/kernel"
)

// WrapIsDislike - 盤を束縛変数として与えます
func WrapIsDislike(position *e.Position) *func(e.Stone, e.Point) bool {
	// 「手番の勝ちなら1、引き分けなら0、手番の負けなら-1を返す関数（自分視点）」を作成します
	// * `color` - 石の色
	var isDislike = func(color e.Stone, z e.Point) bool {
		// 座標取得
		// 432
		// 5S1
		// 678
		var eastZ = z + e.Cell_Dir4[e.Cell_East]
		var northEastZ = z + e.Cell_Dir4[e.Cell_North] + 1
		var northZ = z + e.Cell_Dir4[e.Cell_North]
		var northWestZ = z + e.Cell_Dir4[e.Cell_North] - 1
		var westZ = z + e.Cell_Dir4[e.Cell_West]
		var southWestZ = z + e.Cell_Dir4[e.Cell_South] - 1
		var southZ = z + e.Cell_Dir4[e.Cell_South]
		var southEastZ = z + e.Cell_Dir4[e.Cell_South] + 1

		// 東北
		// **
		// S*
		if isEmptyTriangle(position, color, [3]e.Point{eastZ, northEastZ, northZ}) {
			return true
		}

		// 西北
		// **
		// *S
		if isEmptyTriangle(position, color, [3]e.Point{northZ, northWestZ, westZ}) {
			return true
		}

		// 西南
		// *S
		// **
		if isEmptyTriangle(position, color, [3]e.Point{westZ, southWestZ, southZ}) {
			return true
		}

		// 東南
		// S*
		// **
		if isEmptyTriangle(position, color, [3]e.Point{southZ, southEastZ, eastZ}) {
			return true
		}

		return false
	}

	return &isDislike
}

func isEmptyTriangle(position *e.Position, myColor e.Stone, points [3]e.Point) bool {
	var myColorNum = 0
	var emptyNum = 0

	for _, z := range points {
		var color = position.GetBoard().GetStoneAt(z)
		if color == myColor {
			myColorNum++
		} else if color == e.Stone_Space {
			emptyNum++
		}
	}

	return myColorNum == 2 && emptyNum == 1
}
