package kernel

type Kernel struct {
	// Position - 局面
	Position *Position
	// Record - 棋譜
	Record Record
}

// NewKernel - 新規
func NewKernel(gameRule *GameRule, boardSize int) *Kernel {
	var k = new(Kernel)

	k.Position = NewPosition(gameRule, boardSize)
	k.Record = *NewRecord(gameRule.GetMaxPositionNumber())

	return k
}

// ResizeBoard - 盤サイズの変更
// - 別途、盤面の初期化を行ってください
func (k *Kernel) ResizeBoard(boardSize int) {
	k.Position = NewPosition(&k.Position.board.gameRule, boardSize)
}

// SetPlaceKo - 現局面のコウの番地を設定
func (k *Kernel) SetPlaceKoOfCurrentPosition(placeKo Point) {
	k.Position.KoZ = placeKo
}

// SetPlaceKo - 現局面のコウの番地を設定
func (k *Kernel) GetPlaceKoOfCurrentPosition() Point {
	return k.Position.KoZ
}
