package main

import (
	"time"

	code "github.com/muzudho/kifuwarabe-go-toramedal/engine/coding_obj"
	e "github.com/muzudho/kifuwarabe-go-toramedal/engine/kernel"
	pl "github.com/muzudho/kifuwarabe-go-toramedal/engine/play_algorithm"
	p "github.com/muzudho/kifuwarabe-go-toramedal/engine/presenter"
)

// SelfPlay - コンピューター同士の対局。
func SelfPlay(kernel *e.Kernel) {
	code.Console.Trace("# GoGo SelfPlay 自己対局開始☆（＾～＾）\n")

	var color = e.Stone_Black

	for {
		var z = GetComputerMoveDuringSelfPlay(kernel, color)

		var recItem = new(e.RecordItem)
		recItem.Z = z
		e.PutStoneOnRecord(kernel, z, color, recItem)

		code.Console.Print("z=%s,color=%d", kernel.Position.GetBoard().GetCoordinate().GetGtpMoveFromPoint(z), color) // テスト
		// p.PrintCheckBoard(position)                                        // テスト
		p.PrintBoard(kernel, kernel.Position.Number)

		// パスで２手目以降で棋譜の１つ前（相手）もパスなら終了します。
		if z == e.Cell_Pass && 1 < kernel.Position.Number && kernel.Record[kernel.Position.Number-2].GetZ() == e.Cell_Pass {
			break
		}

		// 手数上限に至ったら抜ける
		if kernel.Position.GetBoard().GetGameRule().GetMaxPositionNumber() <= kernel.Position.Number {
			break
		}

		color = e.FlipColor(color)
	}

	p.PrintSgf(kernel, kernel.Position.Number, kernel.Record)
}

// GetComputerMoveDuringSelfPlay - コンピューターの指し手。 SelfplayLesson09 から呼び出されます
func GetComputerMoveDuringSelfPlay(kernel *e.Kernel, color e.Stone) e.Point {

	var start = time.Now()
	pl.AllPlayouts = 0

	var z, winRate = pl.GetBestZByUct(
		kernel,
		color,
		createPrintingOfCalc(kernel),
		createPrintingOfCalcFin(kernel))

	var sec = time.Since(start).Seconds()
	code.Console.Info("(GetComputerMoveDuringSelfPlay) %.1f sec, %.0f playout/sec, play_z=%04d,rate=%.4f,positionNumber=%d,color=%d,playouts=%d\n",
		sec, float64(pl.AllPlayouts)/sec, kernel.Position.GetBoard().GetCoordinate().GetZ4FromPoint(z), winRate, kernel.Position.Number, color, pl.AllPlayouts)
	return z
}
