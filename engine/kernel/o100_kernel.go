package kernel

type Kernel struct {
	// 局面
	Position *Position
}

func NewKernel() *Kernel {
	var k = new(Kernel)

	// 19路盤で新規作成
	k.Position = NewPosition(19)

	return k
}
