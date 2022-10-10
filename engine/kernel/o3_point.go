package kernel

// Point - 交点の座標。壁を含む盤の左上を 0 とします
type Point int

// Cell_Pass - パス
const Cell_Pass Point = 0

// IllegalZ - 石が置けない番地の目印として使用。例：UCT計算中に石が置けなかった
const Cell_Illegal Point = -1

type Cell_Direction int

// Cell_Dir4 配列のインデックスに対応
const (
	Cell_East Cell_Direction = iota
	Cell_North
	Cell_West
	Cell_South
)
