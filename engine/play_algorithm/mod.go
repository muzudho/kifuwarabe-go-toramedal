package play_algorithm

import (
	e "github.com/muzudho/kifuwarabe-go-toramedal/engine/kernel"
	gd "github.com/muzudho/kifuwarabe-go-toramedal/game_domain"
)

// AllPlayouts - プレイアウトした回数。
var AllPlayouts int

var GettingOfWinnerOnDuringUCTPlayout *func(e.Stone) int
var IsDislike *func(e.Stone, e.Point) bool

// FlagTestPlayout - ？。
var FlagTestPlayout int

func InitKernel(kernel *e.Kernel) {
	// 盤サイズが変わっていることもあるので、先に初期化します
	kernel.InitPosition()

	GettingOfWinnerOnDuringUCTPlayout = WrapGettingOfWinner(kernel.Position)
	IsDislike = gd.WrapIsDislike(kernel.Position)
	AdjustParameters(kernel)
}
