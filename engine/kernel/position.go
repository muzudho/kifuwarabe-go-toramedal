package kernel

// Position - 盤
type Position struct {
	// 盤
	board *Board
	// チェック盤。呼吸点を数えるのに使う
	checkBoard *CheckBoard
	// KoZ - コウの交点。Idx（配列のインデックス）表示。 0 ならコウは無し？
	KoZ Point
	// MovesNum - 手数
	MovesNum int
	// Record - 棋譜
	Record []*RecordItem
	// 二重ループ
	foreachPointWithoutWall func(func(Point))
	// UCT計算アルゴリズム
	uctAlgorithm *UctAlgorithm
}

// NewPosition - 空っぽの局面を生成します
// あとで InitPosition() を呼び出してください
func NewPosition(boardSize int) *Position {
	var p = new(Position)

	p.board = NewBoard()
	p.checkBoard = NewCheckBoard((boardSize + 2) ^ 2)
	p.uctAlgorithm = NewUctAlgorithm(p.board)

	return p
}

// GetBoard - 盤取得
func (p *Position) GetBoard() *Board {
	return p.board
}

// GetCheckBoard - チェック盤取得
func (p *Position) GetCheckBoard() *CheckBoard {
	return p.checkBoard
}

// GetUctAlgorithm - UCT算法
func (p *Position) GetUctAlgorithm() *UctAlgorithm {
	return p.uctAlgorithm
}

// TemporaryPosition - 盤をコピーするときの一時メモリーとして使います
type TemporaryPosition struct {
	// 盤
	Board []Stone
	// KoZ - コウの交点。Idx（配列のインデックス）表示。 0 ならコウは無し？
	KoZ Point
}

// CopyPosition - 盤データのコピー。
func (k *Kernel) CopyPosition() *TemporaryPosition {
	var temp = new(TemporaryPosition)
	temp.Board = make([]Stone, k.BoardCoordinate.GetMemoryBoardArea())
	copy(temp.Board[:], k.Position.board.GetSlice())
	temp.KoZ = k.Position.KoZ
	return temp
}

// ImportPosition - 盤データのコピー。
func (position *Position) ImportPosition(temp *TemporaryPosition) {
	copy(position.board.GetSlice(), temp.Board[:])
	position.KoZ = temp.KoZ
}

// InitPosition - 局面の初期化。
func (k *Kernel) InitPosition() {
	k.Position.Record = make([]*RecordItem, MaxPositionNumber)
	k.Position.uctAlgorithm.uctChildrenSize = k.BoardCoordinate.GetBoardArea() + 1

	// サイズが変わっているケースに対応するため、配列の作り直し
	var memoryBoardArea = k.BoardCoordinate.GetMemoryBoardArea()
	k.Position.board.SetCells(make([]Stone, memoryBoardArea))

	k.Position.checkBoard = NewCheckBoard(memoryBoardArea)

	k.Position.foreachPointWithoutWall = k.BoardCoordinate.PackForeachPointWithoutWall()
	Cell_Dir4 = [4]Point{1, Point(-k.BoardCoordinate.GetMemoryBoardWidth()), -1, Point(k.BoardCoordinate.GetMemoryBoardWidth())}

	// 壁枠を設定
	k.Position.board.DrawWall(memoryBoardArea)

	// 盤上の石を全部消します
	k.Position.board.EraseBoard(k.Position.foreachPointWithoutWall)

	k.Position.MovesNum = 0
	k.Position.KoZ = 0 // コウの指定がないので消します
}

// IterateWithoutWall - 盤イテレーター
func (position *Position) IterateWithoutWall(onPoint func(Point)) {
	position.foreachPointWithoutWall(onPoint)
}
