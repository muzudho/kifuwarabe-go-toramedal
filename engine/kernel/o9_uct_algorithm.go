package kernel

type UctAlgorithm struct {
	// 盤
	board *Board
	// UCT計算中の子の数
	uctChildrenSize int
}

// NewUctAlgorithm - 新規作成
func NewUctAlgorithm(board *Board) *UctAlgorithm {
	var u = new(UctAlgorithm)

	u.board = board

	return u
}

// GetBoard - 盤取得
func (u *UctAlgorithm) GetBoard() *Board {
	return u.board
}

// UctChildrenSize - UCTの最大手数
func (u *UctAlgorithm) UctChildrenSize() int {
	return u.uctChildrenSize
}
