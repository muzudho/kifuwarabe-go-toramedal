package kernel

type UctAlgorithm struct {
	// UCT計算中の子の数
	uctChildrenSize int
}

func NewUctAlgorithm() *UctAlgorithm {
	var u = new(UctAlgorithm)
	return u
}
