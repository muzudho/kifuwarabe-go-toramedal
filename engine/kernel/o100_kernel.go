package kernel

type Kernel struct {
	// 局面
	Position *Position
}

func NewKernel() *Kernel {
	var k = new(Kernel)

	k.Position = NewPosition()

	return k
}
