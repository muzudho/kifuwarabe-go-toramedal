package kernel

import (
	"fmt"
	"strings"
)

// 片方の枠の厚み。東、北、西、南
// const oneSideWallThickness = 1

// 両側の枠の厚み。南北、または東西
const bothSidesWallThickness = 2

// Cell_4Directions - 東、北、西、南を指す配列のインデックスに対応
type Cell_4Directions int

// 東、北、西、南を指す配列のインデックスに対応
const (
	Cell_East Cell_4Directions = iota
	Cell_North
	Cell_West
	Cell_South
)

// BoardCoordinate - 盤座標
type BoardCoordinate struct {
	// 枠付きの盤の水平一辺の交点の要素数
	memoryWidth int
	// 枠付きの盤の垂直一辺の交点の要素数
	memoryHeight int

	// ４方向（東、北、西、南）への相対番地。2015年講習会サンプル、GoGo とは順序が違います
	cell4Directions [4]Point
}

// GetMemoryWidth - 枠付きの盤の水平一辺の交点数
func (bc *BoardCoordinate) GetMemoryWidth() int {
	return bc.memoryWidth
}

// GetMemoryHeight - 枠付きの盤の垂直一辺の交点数
func (bc *BoardCoordinate) GetMemoryHeight() int {
	return bc.memoryHeight
}

// GetMemoryArea - 枠付き盤の面積
func (bc *BoardCoordinate) GetMemoryArea() int {
	return bc.GetMemoryWidth() * bc.GetMemoryHeight()
}

func (bc *BoardCoordinate) GetWidth() int {
	// 枠の分、２つ減らす
	return bc.memoryWidth - bothSidesWallThickness
}

// GetHeight - 枠無し盤の縦幅
func (bc *BoardCoordinate) GetHeight() int {
	// 枠の分、２つ減らす
	return bc.memoryHeight - bothSidesWallThickness
}

// GetBoardArea - 枠無し盤の面積
func (bc *BoardCoordinate) GetBoardArea() int {
	return bc.GetWidth() * bc.GetHeight()
}

// GetRelativePointOf - ４方向（東、北、西、南）の先の番地
func (bc *BoardCoordinate) GetRelativePointOf(dir4 Cell_4Directions) Point {
	return bc.cell4Directions[dir4]
}

// GetEastOf - 東
func (bc *BoardCoordinate) GetEastOf(point Point) Point {
	return point + bc.cell4Directions[Cell_East]
}

// GetNorthEastOf - 北東
func (bc *BoardCoordinate) GetNorthEastOf(point Point) Point {
	return point + bc.cell4Directions[Cell_North] + bc.cell4Directions[Cell_East]
}

// GetNorthOf - 北
func (bc *BoardCoordinate) GetNorthOf(point Point) Point {
	return point + bc.cell4Directions[Cell_North]
}

// GetNorthWestOf - 北西
func (bc *BoardCoordinate) GetNorthWestOf(point Point) Point {
	return point + bc.cell4Directions[Cell_North] + bc.cell4Directions[Cell_West]
}

// GetWestOf - 西
func (bc *BoardCoordinate) GetWestOf(point Point) Point {
	return point + bc.cell4Directions[Cell_West]
}

// GetSouthWestOf - 南西
func (bc *BoardCoordinate) GetSouthWestOf(point Point) Point {
	return point + bc.cell4Directions[Cell_South] + bc.cell4Directions[Cell_West]
}

// GetSouthOf - 南
func (bc *BoardCoordinate) GetSouthOf(point Point) Point {
	return point + bc.cell4Directions[Cell_South]
}

// GetSouthEastOf - 南東
func (bc *BoardCoordinate) GetSouthEastOf(point Point) Point {
	return point + bc.cell4Directions[Cell_South] + bc.cell4Directions[Cell_East]
}

// GetPointFromXy - x,y 形式の座標を、 point （配列のインデックス）へ変換します。
// point は枠を含む盤上での座標です
//
// Parameters
// ----------
// x : int
// 枠を含む盤での筋番号。 Example: 19路盤なら0～20
// y : int
// 枠を含む盤での段番号。 Example: 19路盤なら0～20
//
// Returns
// -------
// point : Point
// 配列インデックス
func (bc *BoardCoordinate) GetPointFromXy(x int, y int) Point {
	return Point(y*bc.GetMemoryWidth() + x)
}

// GetPointFromGtpMove - GTPの座標符号を Point に変換します
// * `gtp_move` - 最初の１文字はアルファベット、２文字目（あれば３文字目）は数字と想定。 例: q10
func (bc *BoardCoordinate) GetPointFromGtpMove(gtp_move string) Point {
	gtp_move = strings.ToUpper(gtp_move)

	if gtp_move == "PASS" {
		return 0
	}

	// 筋 0～
	var x = gtp_move[0] - 'A' + 1
	if gtp_move[0] >= 'I' {
		x--
	}

	// 段 0～
	var y = int(gtp_move[1] - '0')
	if 2 < len(gtp_move) {
		y *= 10
		y += int(gtp_move[2] - '0')
	}

	// インデックス
	var z = bc.GetPointFromXy(int(x), y)
	return z
}

// GetGtpMoveFromPoint - 番地をGTP用の指し手に変換。 例: Q10
func (bc *BoardCoordinate) GetGtpMoveFromPoint(point Point) string {
	if point == 0 {
		return "PASS"
	} else if point == Cell_Illegal {
		return "PASS" // 仕方なく
		// return "ILLEGAL" // GTP の仕様外です
	}

	var y = int(point) / bc.GetMemoryWidth()
	var x = int(point) % bc.GetMemoryWidth()

	// 筋が25（'Z'）より大きくなることは想定していません
	var alphabet_x = 'A' + x - 1
	if alphabet_x >= 'I' {
		alphabet_x++
	}

	// code.Console.Debug("y=%d x=%d z=%d alphabet_x=%d alphabet_x(char)=%c\n", y, x, z, alphabet_x, alphabet_x)

	return fmt.Sprintf("%c%d", alphabet_x, y)
}

// ForeachCellWithoutWall -  盤の（枠を除く）全ての交点に順にアクセスします
func (bc *BoardCoordinate) ForeachCellWithoutWall(setPoint func(Point)) {
	// - x,y は枠無しの盤での0から始まる座標
	// - point は枠無しの盤を配列にしたもので0から始まる配列の添え字
	var heightNth = bc.memoryHeight - 1
	var widthNth = bc.memoryWidth - 1
	for y := 1; y < heightNth; y++ {
		for x := 1; x < widthNth; x++ {
			var point = bc.GetPointFromXy(x, y)
			setPoint(point)
		}
	}
}

// GetZ4FromPoint - z（配列のインデックス）を XXYY形式（3～4桁の数）の座標へ変換します。
func (bc *BoardCoordinate) GetZ4FromPoint(point Point) int {
	if point == 0 {
		return 0
	}
	var y = int(point) / bc.GetMemoryWidth()
	var x = int(point) - y*bc.GetMemoryWidth()
	return x*100 + y
}
