package kernel

// BoardCoordinate - 盤座標
type BoardCoordinate struct {
	// 枠付きの盤の一辺の交点の要素数
	memoryWidth int
}

func NewBoardCoordinate(boardSize int) *BoardCoordinate {
	var bc = new(BoardCoordinate)

	// 枠の分、２つ増える
	bc.memoryWidth = boardSize + 2

	return bc
}

// BoardSize - 何路盤
var BoardSize int

func (bc *BoardCoordinate) SetBoardSize(boardSize int) {
	// 枠の分、２つ増える
	bc.memoryWidth = boardSize + 2
}
func SetBoardSize(boardSize int) {
	BoardSize = boardSize
}

func (bc *BoardCoordinate) GetBoardWidth() int {
	// 枠の分、２つ減らす
	return bc.memoryWidth - 2
}

// GetBoardArea - 壁無し盤の面積
func (bc *BoardCoordinate) GetBoardArea() int {
	return bc.GetBoardWidth() * bc.GetBoardWidth()
}

// GetMemoryBoardWidth - 枠付きの盤の一辺の交点数
func (bc *BoardCoordinate) GetMemoryBoardWidth() int {
	return bc.memoryWidth
}

// GetMemoryBoardWidth - 枠付きの盤の一辺の交点数
func GetMemoryBoardWidth() int {
	return BoardSize + 2
}

// GetMemoryBoardArea - 壁付き盤の面積
func (bc *BoardCoordinate) GetMemoryBoardArea() int {
	return bc.GetMemoryBoardWidth() * bc.GetMemoryBoardWidth()
}

// GetZ4FromPoint - z（配列のインデックス）を XXYY形式（3～4桁の数）の座標へ変換します。
func (bc *BoardCoordinate) GetZ4FromPoint(point Point) int {
	if point == 0 {
		return 0
	}
	var y = int(point) / bc.GetMemoryBoardWidth()
	var x = int(point) - y*bc.GetMemoryBoardWidth()
	return x*100 + y
}
