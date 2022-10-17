package kernel

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
	// 着手点を含む連
	placePlayRen Ren
}

// searchStoneLibInfo - 呼吸点の計算
func searchStoneLibInfo(k *Kernel, here Point, color Stone, ls *LibertySearchAlgorithm) LibertyInfo {
	var oppColor = FlipColor(color) //相手(opponent)の石の色
	var libInfo = LibertyInfo{0, 0, 0, 0, [4]*Ren{nil, nil, nil, nil}, Ren{0, 0, color}}

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

		ls.CountLiberty(p, &libInfo.placePlayRen)

		libInfo.around[dir].LibertyArea = libInfo.placePlayRen.LibertyArea // 呼吸点の数
		libInfo.around[dir].StoneArea = libInfo.placePlayRen.StoneArea     // 連の意地の数
		libInfo.around[dir].Color = stone                                  // 石の色

		if stone == oppColor && libInfo.placePlayRen.LibertyArea == 1 { // 相手の石で、呼吸点が１つで、その呼吸点に今石を置いたなら
			libInfo.captureSum += libInfo.placePlayRen.StoneArea
		}
		if stone == color && 2 <= libInfo.placePlayRen.LibertyArea { // 隣接する連が自分の石で、その石が呼吸点を２つ持ってるようなら
			libInfo.myBreathFriend++
		}

		// TODO ここらへんで再帰したいが
	}
	k.Position.board.ForeachNeumannNeighborhood(here, eachAdjacent)

	return libInfo
}

// PutStone - 石を置きます。
// * `placePlay` - 着手点
//
// # Returns
// エラーコード
func PutStone(k *Kernel, placePlay Point, color Stone) int {

	if placePlay == Cell_Pass { // 投了なら、コウを消して関数を正常終了
		k.ClearPlaceKoOfCurrentPosition()
		return 0
	}

	var oppColor = FlipColor(color) //相手(opponent)の石の色
	var ls = NewLibertySearchAlgorithm(&k.Position.board, &k.Position.checkBoard)

	// 呼吸点を計算します
	var libInfo = searchStoneLibInfo(k, placePlay, color, ls)

	// 石を置くと明らかに損なケース、また、ルール上石を置けないケースなら、石を置きません
	if libInfo.captureSum == 0 && libInfo.space == 0 && libInfo.myBreathFriend == 0 {
		// 例えば黒番で 1 の箇所に打つのは損なので、石を置きません
		//
		//  ooo
		// ox1o
		//  oxo
		//   o
		return 1
	}
	if k.IsPutStoneOnKo(placePlay) { // コウに石を置こうとしたか？
		return 2
	}
	if libInfo.wall+libInfo.myBreathFriend == 4 {
		// 例えば黒番で 1, 2 の箇所（眼）に打つのは損なので、石を置きません
		//
		// #########
		//  x2x  x1#
		//   x    x#
		//         #
		return 3
	}

	if k.Position.GetBoard().IsMasonry(placePlay) { // 石の上に石を置こうとしたか
		return 4
	}

	k.ClearPlaceKoOfCurrentPosition() // コウを消します

	// 石を取り上げます
	var eachAdjacentPointWhenCaptureStone = func(dir Cell_4Directions, p Point) {
		var lib = libInfo.around[dir].LibertyArea // 隣接する連の呼吸点の数
		var adjColor = libInfo.around[dir].Color  // 隣接する連の石の色

		if adjColor == oppColor && // 隣接する連が相手の石で（枠はここで除外されます）
			lib == 1 && // その呼吸点は１つで、そこに今石を置いた
			!k.Position.GetBoard().IsSpaceAt(p) { // 石はまだあるなら（上と右の石がくっついている、といったことを除外）

			k.Position.GetBoard().FillRen(p, oppColor)

			// もし取った石の数が１個なら、その石のある隣接した交点はコウ。また、図形上、コウは１個しか出現しません
			if libInfo.around[dir].StoneArea == 1 {
				k.SetPlaceKoOfCurrentPosition(p)
			}
		}
	}
	k.Position.board.ForeachNeumannNeighborhood(placePlay, eachAdjacentPointWhenCaptureStone)

	k.Position.GetBoard().SetStoneAt(placePlay, color)
	ls.CountLiberty(placePlay, &libInfo.placePlayRen)

	return 0
}
