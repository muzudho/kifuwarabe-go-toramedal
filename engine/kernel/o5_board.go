package kernel

type Board struct {
	// 盤座標
	coordinate *BoardCoordinate

	// 盤
	cells []Stone
}

func NewBoard(boardSize int) *Board {
	var b = new(Board)

	b.coordinate = NewBoardCoordinate(boardSize)

	return b
}

// GetCoordinate - 盤座標取得
func (b *Board) GetCoordinate() *BoardCoordinate {
	return b.coordinate
}

// GetSlice - 配列のスライスを取得
func (b *Board) GetSlice() []Stone {
	return b.cells[:]
}

// GetStoneAt - 交点を指定して石の色を取得
func (b *Board) GetStoneAt(point Point) Stone {
	return b.cells[point]
}

// SetStoneAt - 交点を指定して石の色を設定
func (b *Board) SetStoneAt(point Point, stone Stone) {
	b.cells[point] = stone
}

// SetCells - 盤面の設定
func (b *Board) SetCells(cells []Stone) {
	// Go言語での配列の代入は値渡しなのでこれでOK。C言語のようなポインター渡しではない
	b.cells = cells
}

// IsEmpty - 指定の交点は空点か？
func (b *Board) IsSpaceAt(point Point) bool {
	return b.GetStoneAt(point) == Stone_Space
}

// DrawEmptyBoard - 空っぽの盤にします
func (b *Board) DrawEmptyBoard() {
	// 壁枠を設定
	b.DrawWall()

	// 盤上の石を全部消します
	b.EraseBoard()
}

// DrawWall - 壁枠を設定します
func (b *Board) DrawWall() {
	for z := Point(0); z < Point(b.coordinate.GetMemoryBoardArea()); z++ {
		b.SetStoneAt(z, Stone_Wall)
	}
}

// EraseBoard - （壁枠を除く）盤上をすべて空点にします
func (b *Board) EraseBoard() {
	var setPoint = func(point Point) {
		b.SetStoneAt(point, Stone_Space)
	}

	b.coordinate.ForeachPointWithoutWall(setPoint)
}

// FillRen - 石を打ち上げ（取り上げ、取り除き）ます。
func (b *Board) FillRen(z Point, color Stone) {
	b.SetStoneAt(z, Stone_Space) // 石を消します

	for dir := 0; dir < 4; dir++ {
		var adjZ = z + b.coordinate.cellDir4[dir]

		if b.GetStoneAt(adjZ) == color { // 再帰します
			b.FillRen(adjZ, color)
		}
	}
}
