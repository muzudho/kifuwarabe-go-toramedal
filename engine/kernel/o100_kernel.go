package kernel

type Kernel struct {
	// 盤座標
	BoardCoordinate *BoardCoordinate

	// 局面
	Position *Position
}

func NewKernel() *Kernel {
	var k = new(Kernel)

	// 既定値で新規作成
	k.BoardCoordinate = NewBoardCoordinate(19)
	k.Position = NewPosition(19)

	return k
}
