package kernel

type Board struct {
	// 盤
	cells []Stone
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
