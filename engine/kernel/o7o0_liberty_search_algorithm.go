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
func (ls *LibertySearchAlgorithm) CountLiberty(z Point, placePlayRen *Ren) {
	placePlayRen.LibertyArea = 0
	placePlayRen.StoneArea = 0

	// チェックボードの初期化
	var eachPoint = func(z Point) {
		ls.checkBoard.ClearAllBitsAt(z)
	}
	ls.board.GetCoordinate().ForeachCellWithoutWall(eachPoint)

	ls.searchStoneRenRecursive(z, ls.board.GetStoneAt(z), placePlayRen)
}

// 石の連の探索
//
// * 再帰関数
// * `libertyArea` - 呼吸点の数
// * `renArea` - 連の石の数
func (ls *LibertySearchAlgorithm) searchStoneRenRecursive(here Point, color Stone, placePlayRen *Ren) {

	// 石のチェック
	ls.checkBoard.Overwrite(here, Mark_BitStone)

	placePlayRen.StoneArea++

	var eachAdjacent = func(dir Cell_4Directions, p Point) {
		if !ls.checkBoard.IsZeroAt(p) {
			return // あとの処理をスキップ
		}

		if ls.board.IsSpaceAt(p) { // 空点

			ls.checkBoard.Overwrite(p, Mark_BitStone)

			placePlayRen.StoneArea++
		} else if ls.board.GetStoneAt(p) == color {
			ls.searchStoneRenRecursive(p, color, placePlayRen) // 再帰
		}
	}

	// 隣接する４方向
	ls.board.ForeachNeumannNeighborhood(here, eachAdjacent)
}
