package kernel

// Stone - 石の色
type Stone int

const (
	// Stone_Space - 空点
	Stone_Space Stone = iota
	// Stone_Black - 黒石
	Stone_Black
	// Stone_White - 白石
	Stone_White
	// Stone_Wall - 壁
	Stone_Wall
)
