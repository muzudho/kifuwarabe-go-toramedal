package kernel

// LibertySearchAlgorithm - 呼吸点探索アルゴリズム
type LibertySearchAlgorithm struct {
	// 盤
	board *Board
	// チェック盤
	checkBoard *CheckBoard
}

// NewLibertySearchAlgorithm - 新規作成
func NewLibertySearchAlgorithm(board *Board, checkBoard *CheckBoard) *LibertySearchAlgorithm {
	var ls = new(LibertySearchAlgorithm)

	ls.board = board
	ls.checkBoard = checkBoard

	return ls
}

// CountLiberty - 呼吸点を数えます。
// * `libertyArea` - 呼吸点の数
// * `renArea` - 連の石の数
func (ls *LibertySearchAlgorithm) CountLiberty(z Point, libertyArea *int, renArea *int) {
	*libertyArea = 0
	*renArea = 0

	// チェックボードの初期化
	var setPoint = func(z Point) {
		ls.checkBoard.ClearAllBitsAt(z)
	}
	ls.board.GetCoordinate().ForeachCellWithoutWall(setPoint)

	ls.countLibertySub(z, ls.board.GetStoneAt(z), libertyArea, renArea)
}

// * `libertyArea` - 呼吸点の数
// * `renArea` - 連の石の数
func (ls *LibertySearchAlgorithm) countLibertySub(z Point, color Stone, libertyArea *int, renArea *int) {

	ls.checkBoard.Overwrite(z, Mark_BitStone)

	*renArea++
	for i := 0; i < 4; i++ {
		var adjZ = z + ls.board.coordinate.cell4Directions[i]

		if !ls.checkBoard.IsZeroAt(adjZ) {
			continue
		}

		if ls.board.IsSpaceAt(adjZ) { // 空点

			ls.checkBoard.Overwrite(adjZ, Mark_BitStone)

			*libertyArea++
		} else if ls.board.GetStoneAt(adjZ) == color {
			ls.countLibertySub(adjZ, color, libertyArea, renArea) // 再帰
		}
	}
}
