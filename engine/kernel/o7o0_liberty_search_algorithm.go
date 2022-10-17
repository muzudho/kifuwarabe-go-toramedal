package kernel

// LibertySearchAlgorithm - 呼吸点探索アルゴリズム
type LibertySearchAlgorithm struct {
	// 盤
	board *Board
	// チェック盤
	checkBoard *CheckBoard
	// 着手点を含む連
	foundRen Ren
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
func (ls *LibertySearchAlgorithm) CountLiberty(z Point) {
	ls.foundRen.LibertyArea = 0
	ls.foundRen.StoneArea = 0

	// チェックボードの初期化
	var eachPoint = func(z Point) {
		ls.checkBoard.ClearAllBitsAt(z)
	}
	ls.board.GetCoordinate().ForeachCellWithoutWall(eachPoint)

	ls.searchStoneRenRecursive(z, ls.board.GetStoneAt(z))
}

// 石の連の探索
//
// * 再帰関数
// * `libertyArea` - 呼吸点の数
// * `renArea` - 連の石の数
func (ls *LibertySearchAlgorithm) searchStoneRenRecursive(here Point, color Stone) {

	// 石のチェック
	ls.checkBoard.Overwrite(here, Mark_BitStone)

	ls.foundRen.StoneArea++

	var eachAdjacent = func(dir Cell_4Directions, p Point) {
		if !ls.checkBoard.IsZeroAt(p) {
			return // あとの処理をスキップ
		}

		if ls.board.IsSpaceAt(p) { // 空点

			ls.checkBoard.Overwrite(p, Mark_BitStone)

			ls.foundRen.StoneArea++
		} else if ls.board.GetStoneAt(p) == color {
			ls.searchStoneRenRecursive(p, color) // 再帰
		}
	}

	// 隣接する４方向
	ls.board.ForeachNeumannNeighborhood(here, eachAdjacent)
}

type LibertyInfo struct {
	// アゲハマの数
	captureSum int
	// 隣接している空点への向きの数
	space int
	// 呼吸できる自分の石と隣接している向きの数
	myBreathFriend int
	// 隣接している枠への向きの数
	wall int
	// 隣接する４つの交点
	around [4]*Ren
}

// searchStoneLibInfo - 呼吸点の計算
func (ls *LibertySearchAlgorithm) searchStoneLibInfo(k *Kernel, here Point, color Stone) LibertyInfo {
	var oppColor = FlipColor(color) //相手(opponent)の石の色
	ls.foundRen = Ren{0, 0, color}
	var libInfo = LibertyInfo{0, 0, 0, 0, [4]*Ren{nil, nil, nil, nil}}

	// 隣接する交点毎に
	var eachAdjacent = func(dir Cell_4Directions, p Point) {
		libInfo.around[dir] = NewRen(0, 0, 0) // 呼吸点の数, 連の石の数, 石の色

		var stone = k.Position.GetBoard().GetStoneAt(p) // 石の色
		switch stone {

		case Stone_Space: // 空点
			libInfo.space++
			return

		case Stone_Wall: // 枠
			libInfo.wall++
			return
		}

		ls.CountLiberty(p)

		libInfo.around[dir].LibertyArea = ls.foundRen.LibertyArea // 呼吸点の数
		libInfo.around[dir].StoneArea = ls.foundRen.StoneArea     // 連の意地の数
		libInfo.around[dir].Color = stone                         // 石の色

		if stone == oppColor && ls.foundRen.LibertyArea == 1 { // 相手の石で、呼吸点が１つで、その呼吸点に今石を置いたなら
			libInfo.captureSum += ls.foundRen.StoneArea
		}
		if stone == color && 2 <= ls.foundRen.LibertyArea { // 隣接する連が自分の石で、その石が呼吸点を２つ持ってるようなら
			libInfo.myBreathFriend++
		}

		// TODO ここらへんで再帰したいが
	}
	k.Position.board.ForeachNeumannNeighborhood(here, eachAdjacent)

	return libInfo
}
