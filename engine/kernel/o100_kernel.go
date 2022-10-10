package kernel

type Kernel struct {
	// 局面
	Position *Position
	// Record - 棋譜
	Record []*RecordItem
}

// NewKernel - 新規
func NewKernel(gameRule *GameRule, boardSize int) *Kernel {
	var k = new(Kernel)

	k.Position = NewPosition(gameRule, boardSize)

	return k
}
