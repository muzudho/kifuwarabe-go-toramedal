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

func SetBoardSize(boardSize int) {
	BoardSize = boardSize
}

// BoardSize - 何路盤
var BoardSize int

// GetBoardArea - 壁無し盤の面積
func GetBoardArea() int {
	return BoardSize * BoardSize
}

// GetMemoryBoardWidth - 枠付きの盤の一辺の交点数
func GetMemoryBoardWidth() int {
	return BoardSize + 2
}

// GetMemoryBoardArea - 壁付き盤の面積
func GetMemoryBoardArea() int {
	return GetMemoryBoardWidth() * GetMemoryBoardWidth()
}
