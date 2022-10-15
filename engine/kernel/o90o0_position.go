package kernel

// Position - 盤
type Position struct {
	// 盤
	board Board
	// チェック盤。呼吸点を数えるのに使う
	checkBoard CheckBoard
	// UCT計算アルゴリズム
	uctAlgorithm UctAlgorithm
}

// NewDirtyPosition - 新規作成するが、初期化されていない
//
// - あとで Kernel#InitPosition() を呼び出してください
func NewDirtyPosition(gameRule GameRule, boardSize int) *Position {
	var p = new(Position)

	p.board = *NewBoard(gameRule, boardSize)
	p.checkBoard = *NewDirtyCheckBoard()
	p.uctAlgorithm = UctAlgorithm{&p.board, 0}

	return p
}

// InitPosition - 局面の初期化。
func (k *Kernel) InitPosition() {
	// 空っぽの盤面に設定
	k.Position.board.DrawEmptyBoard()

	// チェック盤の作り直し
	k.Position.checkBoard.Init(k.Position.board.coordinate)

	// 棋譜の作り直し
	k.Record = *NewRecord(k.Position.board.gameRule.maxPositionNumber, k.Position.board.coordinate.GetMemoryArea())

	// UCTアルゴリズムの初期設定
	k.Position.uctAlgorithm.uctChildrenSize = k.Position.board.coordinate.GetBoardArea() + 1

	k.Record.positionNumber = 0
	k.ClearPlaceKoOfCurrentPosition() // コウの指定がないので消します
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
