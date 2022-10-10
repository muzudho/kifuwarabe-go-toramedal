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
	k.Record = *NewRecord()

	return k
}
