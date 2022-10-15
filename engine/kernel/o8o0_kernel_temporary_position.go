package kernel

// TemporaryPosition - 盤をコピーするときの一時メモリーとして使います
type TemporaryPosition struct {
	// 盤
	Board []Stone
}

// NewCopyPosition - 盤データのコピー。
func (k *Kernel) NewCopyPosition() *TemporaryPosition {
	var temp = new(TemporaryPosition)
	temp.Board = make([]Stone, k.Position.board.coordinate.GetMemoryArea())
	copy(temp.Board[:], k.Position.board.GetSlice())
	return temp
}

// ImportPosition - 盤データのコピー。
func (k *Kernel) ImportPosition(temp *TemporaryPosition) {
	copy(k.Position.board.GetSlice(), temp.Board[:])
}
