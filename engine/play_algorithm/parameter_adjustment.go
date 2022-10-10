package play_algorithm

import (
	e "github.com/muzudho/kifuwarabe-go-toramedal/engine/kernel"
)

// プレイアウトする回数（あとで設定されます）
var PlayoutTrialCount = 0

// UCTをループする回数（あとで設定されます）
var UctLoopCount = 0

// ランダム鳩の巣仮説定数 a。およそ 18
// 面積 * 2 pi e 、つまり およそ 17 で、５００回に１回見落としがある程度、
// 面積 * (2 pi e + 1) 、 つまり およそ 18 で、１万回に１回見落としがある程度の精度（自分調べ）
var randomPigeonA = 17 // 2 * math.Pi * math.E

// ランダム鳩の巣仮説 試行回数 x
// 📖 [random-pigeon-nest-hypothesis](https://github.com/muzudho/random-pigeon-nest-hypothesis)
func GetRandomPigeonX(N int) int {
	return N * randomPigeonA
	// return int(math.Ceil(float64(N) * randomPigeonA))
}

func AdjustParameters(kernel *e.Kernel) {
	var boardSize = kernel.Position.GetBoard().GetCoordinate().GetBoardWidth()
	if boardSize < 10 {
		// 10路盤より小さいとき
		PlayoutTrialCount = boardSize*boardSize + 200
	} else {
		// 19路盤を想定。
		// 19路盤なら 19 × 19 が理想だが、時間切れになるので短くします
		var powerUp = 0.01 // 盤の下辺にペンキ塗りを始めてしまう
		// var powerUp = 0.1 // ランダム打ちに見える
		// var powerUp = 0.125 // すぐ打つ
		// var powerUp = 0.25 // 早いけどランダム打ちに見える。負けてる試合を長く見るだけ
		// var powerUp = 0.33 // 3秒で打てるけど弱い？
		// var powerUp = 0.4
		// var powerUp = 0.45 ----> 30分で切れ負け
		// var powerUp = 0.5 // 5秒で打てる ----> 30分で切れ負け
		// var powerUp = 0.66 // 30分で切れ負け
		// var powerUp = 0.75 // 11秒ぐらいかかる
		// var powerUp = 1 // 30秒ぐらいかかる
		PlayoutTrialCount = int(float64(boardSize*boardSize) * powerUp)
	}

	// 盤面全体を１回は選ぶことを、完璧ではありませんが、ある程度の精度でカバーします
	UctLoopCount = GetRandomPigeonX(kernel.Position.GetBoard().GetCoordinate().GetBoardArea())
}
