package kernel

// PutStone - 石を置きます。
// * `placePlay` - 着手点
//
// # Returns
// エラーコード
func PutStone(k *Kernel, placePlay Point, color Stone) int {
	var around = [4]*Ren{}              // 隣接する４つの交点
	var placePlayRen = Ren{0, 0, color} // 着手点を含む連
	var oppColor = FlipColor(color)     //相手(opponent)の石の色
	var space = 0                       // 隣接している空点への向きの数
	var wall = 0                        // 隣接している枠への向きの数
	var myBreathFriend = 0              // 呼吸できる自分の石と隣接している向きの数
	var captureSum = 0                  // アゲハマの数

	if placePlay == Cell_Pass { // 投了なら、コウを消して関数を正常終了
		k.ClearPlaceKoOfCurrentPosition()
		return 0
	}

	var ls = NewLibertySearchAlgorithm(&k.Position.board, &k.Position.checkBoard)

	// 呼吸点を計算します
	var eachAdjacentPointWhenGetLiberty = func(dir Cell_4Directions, p Point) {
		around[dir] = NewRen(0, 0, 0) // 呼吸点の数, 連の石の数, 石の色

		var adjColor = k.Position.GetBoard().GetStoneAt(p) // 隣(adjacent)の交点の石の色
		if adjColor == Stone_Space {                       // 空点
			space++
			return
		}
		if adjColor == Stone_Wall { // 枠
			wall++
			return
		}
		ls.CountLiberty(p, &placePlayRen)
		around[dir].LibertyArea = placePlayRen.LibertyArea         // 呼吸点の数
		around[dir].StoneArea = placePlayRen.StoneArea             // 連の意地の数
		around[dir].Color = adjColor                               // 石の色
		if adjColor == oppColor && placePlayRen.LibertyArea == 1 { // 相手の石で、呼吸点が１つで、その呼吸点に今石を置いたなら
			captureSum += placePlayRen.StoneArea
		}
		if adjColor == color && 2 <= placePlayRen.LibertyArea { // 隣接する連が自分の石で、その石が呼吸点を２つ持ってるようなら
			myBreathFriend++
		}
	}
	k.Position.board.ForeachNeumannNeighborhood(placePlay, eachAdjacentPointWhenGetLiberty)

	// 石を置くと明らかに損なケース、また、ルール上石を置けないケースなら、石を置きません
	if captureSum == 0 && space == 0 && myBreathFriend == 0 {
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
	if wall+myBreathFriend == 4 {
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
		var lib = around[dir].LibertyArea // 隣接する連の呼吸点の数
		var adjColor = around[dir].Color  // 隣接する連の石の色

		if adjColor == oppColor && // 隣接する連が相手の石で（枠はここで除外されます）
			lib == 1 && // その呼吸点は１つで、そこに今石を置いた
			!k.Position.GetBoard().IsSpaceAt(p) { // 石はまだあるなら（上と右の石がくっついている、といったことを除外）

			k.Position.GetBoard().FillRen(p, oppColor)

			// もし取った石の数が１個なら、その石のある隣接した交点はコウ。また、図形上、コウは１個しか出現しません
			if around[dir].StoneArea == 1 {
				k.SetPlaceKoOfCurrentPosition(p)
			}
		}
	}
	k.Position.board.ForeachNeumannNeighborhood(placePlay, eachAdjacentPointWhenCaptureStone)

	k.Position.GetBoard().SetStoneAt(placePlay, color)
	ls.CountLiberty(placePlay, &placePlayRen)

	return 0
}
