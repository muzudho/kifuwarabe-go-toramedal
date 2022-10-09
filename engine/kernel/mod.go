package kernel

const (
	// Author - 囲碁思考エンジンの作者名だぜ☆（＾～＾）
	Author = "Satoshi Takahashi"
)

type Kernel struct {
	// 局面
	Position *Position
}

func NewKernel() *Kernel {
	var k = new(Kernel)

	k.Position = NewPosition()

	return k
}
