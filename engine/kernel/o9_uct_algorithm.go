package kernel

type UctAlgorithm struct {
	// 盤への参照
	pBoard *Board
	// UCT計算中の子の数
	uctChildrenSize int
}

// GetPtrBoard - 盤へのポインター取得
func (u *UctAlgorithm) GetPtrBoard() *Board {
	return u.pBoard
}

// UctChildrenSize - UCTの最大手数
func (u *UctAlgorithm) UctChildrenSize() int {
	return u.uctChildrenSize
}
