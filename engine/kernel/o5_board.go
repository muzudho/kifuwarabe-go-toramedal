package kernel

type Board struct {
	// 盤
	cells []Stone
}

// GetCells - 交点の配列を取得
func (b *Board) GetCells() []Stone {
	return b.cells
}

// GetCells - 交点の配列を設定
func (b *Board) SetCells(cells []Stone) {
	b.cells = cells
}
