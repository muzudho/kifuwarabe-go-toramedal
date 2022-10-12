package kernel

// TemporaryPosition - 盤をコピーするときの一時メモリーとして使います
type TemporaryPosition struct {
	// 盤
	Board []Stone
	// KoZ - コウの交点。Idx（配列のインデックス）表示。 0 ならコウは無し？
	KoZ Point
}

// NewCopyPosition - 盤データのコピー。
func (k *Kernel) NewCopyPosition() *TemporaryPosition {
	var temp = new(TemporaryPosition)
	temp.Board = make([]Stone, k.Position.board.coordinate.GetMemoryBoardArea())
	copy(temp.Board[:], k.Position.board.GetSlice())
	temp.KoZ = k.GetPlaceKoOfCurrentPosition()
	return temp
}

// ImportPosition - 盤データのコピー。
func (k *Kernel) ImportPosition(temp *TemporaryPosition) {
	copy(k.Position.board.GetSlice(), temp.Board[:])
	k.SetPlaceKoOfCurrentPosition(temp.KoZ)
}
