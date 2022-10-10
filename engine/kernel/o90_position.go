package kernel

// Position - 盤
type Position struct {
	// 盤
	board *Board
	// チェック盤。呼吸点を数えるのに使う
	checkBoard *CheckBoard
	// KoZ - コウの交点。Idx（配列のインデックス）表示。 0 ならコウは無し？
	KoZ Point
	// Number - 手数
	Number PositionNumberInt
	// UCT計算アルゴリズム
	uctAlgorithm *UctAlgorithm
}

// NewPosition - 空っぽの局面を生成します
// あとで InitPosition() を呼び出してください
func NewPosition(gameRule *GameRule, boardSize int) *Position {
	var p = new(Position)

	p.board = NewBoard(gameRule, boardSize)
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

// InitPosition - 局面の初期化。
func (k *Kernel) InitPosition() {
	// 空っぽの盤面に設定
	k.Position.board.SetupEmptyBoard()

	// チェック盤の作り直し
	var memoryBoardArea = k.Position.board.coordinate.GetMemoryBoardArea()
	k.Position.checkBoard = NewCheckBoard(memoryBoardArea)

	// 棋譜の作り直し
	k.Record = *NewRecord(k.Position.board.gameRule.maxPositionNumber)

	// UCTアルゴリズムの初期設定
	k.Position.uctAlgorithm.uctChildrenSize = k.Position.board.coordinate.GetBoardArea() + 1

	k.Position.Number = 0
	k.Position.KoZ = 0 // コウの指定がないので消します
}
