// Source: https://github.com/bleu48/GoGo
// 電通大で行われたコンピュータ囲碁講習会をGolangで追う

package main

import (
	"flag"
	"math/rand"
	"time"

	code "github.com/muzudho/kifuwarabe-go-toramedal/engine/coding_obj"
	cnf "github.com/muzudho/kifuwarabe-go-toramedal/engine/config_obj"
	e "github.com/muzudho/kifuwarabe-go-toramedal/engine/kernel"
	pl "github.com/muzudho/kifuwarabe-go-toramedal/engine/play_algorithm"
)

func main() {
	flag.Parse()
	lessonVer := flag.Arg(0)

	// 乱数の種を設定
	rand.Seed(time.Now().UnixNano())

	// ログの書込み先設定
	code.GtpLog.SetPath("output/gtp_print.log")
	code.ConsoleLog.SetPath(
		"output/trace.log",
		"output/debug.log",
		"output/info.log",
		"output/notice.log",
		"output/warn.log",
		"output/error.log",
		"output/fatal.log",
		"output/print.log")

	code.Console.Trace("# Author: %s\n", e.Author)

	// 設定は囲碁GUIから与えられて上書きされる想定です。設定ファイルはデフォルト設定です
	var config = cnf.LoadGameConf("input/game_conf.toml", OnFatal)
	e.Komi = e.KomiFloat(config.Komi())
	e.MaxPositionNumber = e.PositionNumberInt(config.MaxPositionNumber())

	var kernel = e.NewKernel()

	kernel.BoardCoordinate.SetBoardSize(config.BoardSize())
	e.SetBoardSize(config.BoardSize())

	pl.InitKernel(kernel)
	kernel.Position.SetBoard(config.GetBoardArray())

	if lessonVer == "SelfPlay" {
		SelfPlay(kernel)
	} else {
		RunGtpEngine(kernel) // GTP
	}
}

func OnFatal(errorMessage string) {
	code.Console.Fatal(errorMessage)
}

func createPrintingOfCalc(kernel *e.Kernel) *func(*e.Position, int, e.Point, float64, int) {
	// UCT計算中の表示
	var fn = func(position *e.Position, i int, z e.Point, rate float64, games int) {
		code.Console.Info("(UCT Calculating...) %2d:z=%s,rate=%.4f,games=%3d\n", i, kernel.BoardCoordinate.GetGtpMoveFromPoint(z), rate, games)
	}

	return &fn
}

func createPrintingOfCalcFin(kernel *e.Kernel) *func(*e.Position, e.Point, float64, int, int, int) {
	// UCT計算後の表示
	var fn = func(position *e.Position, bestZ e.Point, rate float64, max int, allPlayouts int, nodeNum int) {
		code.Console.Info("(UCT Calculated    ) bestZ=%s,rate=%.4f,games=%d,playouts=%d,nodes=%d\n",
			kernel.BoardCoordinate.GetGtpMoveFromPoint(bestZ), rate, max, allPlayouts, nodeNum)

	}

	return &fn
}
