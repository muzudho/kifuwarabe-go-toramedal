package kernel

import (
	"math/rand"
)

// Position - 盤
type Position struct {
	// 盤
	board *Board
	// 呼吸点を数えるための一時盤
	checkBoard []int
	// KoZ - コウの交点。Idx（配列のインデックス）表示。 0 ならコウは無し？
	KoZ Point
	// MovesNum - 手数
	MovesNum int
	// Record - 棋譜
	Record []*RecordItem
	// 二重ループ
	iteratorWithoutWall func(func(Point))
	// UCT計算中の子の数
	uctChildrenSize int
}

func (p *Position) GetBoard() *Board {
	return p.board
}

// TemporaryPosition - 盤をコピーするときの一時メモリーとして使います
type TemporaryPosition struct {
	// 盤
	Board []Stone
	// KoZ - コウの交点。Idx（配列のインデックス）表示。 0 ならコウは無し？
	KoZ Point
}

// CopyPosition - 盤データのコピー。
func (k *Kernel) CopyPosition() *TemporaryPosition {
	var temp = new(TemporaryPosition)
	temp.Board = make([]Stone, k.BoardCoordinate.GetMemoryBoardArea())
	copy(temp.Board[:], k.Position.board.GetSlice())
	temp.KoZ = k.Position.KoZ
	return temp
}

// ImportPosition - 盤データのコピー。
func (position *Position) ImportPosition(temp *TemporaryPosition) {
	copy(position.board.GetSlice(), temp.Board[:])
	position.KoZ = temp.KoZ
}

// NewPosition - 空っぽの局面を生成します
// あとで InitPosition() を呼び出してください
func NewPosition() *Position {
	return new(Position)
}

// InitPosition - 局面の初期化。
func (k *Kernel) InitPosition() {
	k.Position.Record = make([]*RecordItem, MaxPositionNumber)
	k.Position.uctChildrenSize = k.BoardCoordinate.GetBoardArea() + 1

	// サイズが変わっているケースに対応するため、配列の作り直し
	var boardMax = k.BoardCoordinate.GetMemoryBoardArea()
	k.Position.board.SetCells(make([]Stone, boardMax))
	k.Position.checkBoard = make([]int, boardMax)
	k.Position.iteratorWithoutWall = CreateBoardIteratorWithoutWall(k)
	Cell_Dir4 = [4]Point{1, Point(-k.BoardCoordinate.GetMemoryBoardWidth()), -1, Point(k.BoardCoordinate.GetMemoryBoardWidth())}

	// 枠線
	for z := Point(0); z < Point(boardMax); z++ {
		k.Position.GetBoard().SetStoneAt(z, Stone_Wall)
	}

	// 盤上
	var onPoint = func(z Point) {
		k.Position.GetBoard().SetStoneAt(z, 0)
	}
	k.Position.iteratorWithoutWall(onPoint)

	k.Position.MovesNum = 0
	k.Position.KoZ = 0 // コウの指定がないので消します
}

// CheckAt - 指定した交点のチェック
func (position *Position) CheckAt(z Point) int {
	return position.checkBoard[z]
}

// ColorAtXy - 指定した交点の石の色
func (k *Kernel) ColorAtXy(x int, y int) Stone {
	var point = Point((y+1)*k.BoardCoordinate.GetMemoryBoardWidth() + x + 1)
	return k.Position.board.GetStoneAt(point)
}

// GetEmptyZ - 空点の z （配列のインデックス）を返します。
func (k *Kernel) GetEmptyZ() Point {
	var x, y int
	var z Point
	for {
		// ランダムに交点を選んで、空点を見つけるまで繰り返します。
		x = rand.Intn(9)
		y = rand.Intn(9)
		z = k.BoardCoordinate.GetPointFromXy(x, y)
		if k.Position.GetBoard().IsSpaceAt(z) { // 空点
			break
		}
	}
	return z
}

// CountLiberty - 呼吸点を数えます。
// * `libertyArea` - 呼吸点の数
// * `renArea` - 連の石の数
func (position *Position) CountLiberty(z Point, libertyArea *int, renArea *int) {
	*libertyArea = 0
	*renArea = 0

	// チェックボードの初期化
	var onPoint = func(z Point) {
		position.checkBoard[z] = 0
	}
	position.iteratorWithoutWall(onPoint)

	position.countLibertySub(z, position.board.GetStoneAt(z), libertyArea, renArea)
}

// * `libertyArea` - 呼吸点の数
// * `renArea` - 連の石の数
func (position *Position) countLibertySub(z Point, color Stone, libertyArea *int, renArea *int) {
	position.checkBoard[z] = 1
	*renArea++
	for i := 0; i < 4; i++ {
		var adjZ = z + Cell_Dir4[i]
		if position.checkBoard[adjZ] != 0 {
			continue
		}
		if position.GetBoard().IsSpaceAt(adjZ) { // 空点
			position.checkBoard[adjZ] = 1
			*libertyArea++
		} else if position.board.GetStoneAt(adjZ) == color {
			position.countLibertySub(adjZ, color, libertyArea, renArea) // 再帰
		}
	}
}

// TakeStone - 石を打ち上げ（取り上げ、取り除き）ます。
func (position *Position) TakeStone(z Point, color Stone) {
	position.board.SetStoneAt(z, Stone_Space) // 石を消します

	for dir := 0; dir < 4; dir++ {
		var adjZ = z + Cell_Dir4[dir]

		if position.board.GetStoneAt(adjZ) == color { // 再帰します
			position.TakeStone(adjZ, color)
		}
	}
}

// IterateWithoutWall - 盤イテレーター
func (position *Position) IterateWithoutWall(onPoint func(Point)) {
	position.iteratorWithoutWall(onPoint)
}

// UctChildrenSize - UCTの最大手数
func (position *Position) UctChildrenSize() int {
	return position.uctChildrenSize
}

// CreateBoardIteratorWithoutWall - 盤の（壁を除く）全ての交点に順にアクセスする boardIterator 関数を生成します
func CreateBoardIteratorWithoutWall(kernel *Kernel) func(func(Point)) {

	var boardSize = kernel.BoardCoordinate.GetBoardWidth()
	var boardIterator = func(onPoint func(Point)) {

		// x,y は壁無しの盤での0から始まる座標、 z は壁有りの盤での0から始まる座標
		for y := 0; y < boardSize; y++ {
			for x := 0; x < boardSize; x++ {
				var z = kernel.BoardCoordinate.GetPointFromXy(x, y)
				onPoint(z)
			}
		}
	}

	return boardIterator
}
