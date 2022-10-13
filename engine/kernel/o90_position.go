package kernel

// Position - 盤
type Position struct {
	// 盤
	board Board
	// チェック盤。呼吸点を数えるのに使う
	checkBoard CheckBoard
	// Number - 手数
	Number PositionNumberInt
	// UCT計算アルゴリズム
	uctAlgorithm UctAlgorithm
}

// NewPosition - 空っぽの局面を生成します
// あとで InitPosition() を呼び出してください
func NewPosition(gameRule GameRule, boardSize int) *Position {
	var p = new(Position)

	var memoryBoardSize = boardSize + 2

	p.board = *NewBoard(gameRule, boardSize)
	p.checkBoard = *NewCheckBoard(memoryBoardSize * memoryBoardSize)
	p.uctAlgorithm = UctAlgorithm{&p.board, 0}

	return p
}

// GetBoard - 盤取得
func (p *Position) GetBoard() *Board {
	return &p.board
}

// GetCheckBoard - チェック盤取得
func (p *Position) GetCheckBoard() *CheckBoard {
	return &p.checkBoard
}

// GetPtrUctAlgorithm - UCT算法へのポインター取得
func (p *Position) GetPtrUctAlgorithm() *UctAlgorithm {
	return &p.uctAlgorithm
}

// InitPosition - 局面の初期化。
func (k *Kernel) InitPosition() {
	// 空っぽの盤面に設定
	k.Position.board.DrawEmptyBoard()

	// チェック盤の作り直し
	var memoryBoardArea = k.Position.board.coordinate.GetMemoryArea()
	k.Position.checkBoard = *NewCheckBoard(memoryBoardArea)

	// 棋譜の作り直し
	k.Record = *NewRecord(k.Position.board.gameRule.maxPositionNumber)

	// UCTアルゴリズムの初期設定
	k.Position.uctAlgorithm.uctChildrenSize = k.Position.board.coordinate.GetBoardArea() + 1

	k.Position.Number = 0
	k.ClearPlaceKoOfCurrentPosition() // コウの指定がないので消します
}
