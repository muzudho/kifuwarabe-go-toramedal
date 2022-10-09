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
		e.PutStoneOnRecord(kernel.Position, z, color, recItem)

		code.Console.Print("z=%s,color=%d", e.GetGtpMoveFromPoint(z), color) // テスト
		// p.PrintCheckBoard(position)                                        // テスト
		p.PrintBoard(kernel, kernel.Position.MovesNum)

		// パスで２手目以降で棋譜の１つ前（相手）もパスなら終了します。
		if z == e.Cell_Pass && 1 < kernel.Position.MovesNum && kernel.Position.Record[kernel.Position.MovesNum-2].GetZ() == e.Cell_Pass {
			break
		}
		// 自己対局は400手で終了します。
		if 400 < kernel.Position.MovesNum {
			break
		} // too long
		color = e.FlipColor(color)
	}

	p.PrintSgf(kernel, kernel.Position.MovesNum, kernel.Position.Record)
}

// GetComputerMoveDuringSelfPlay - コンピューターの指し手。 SelfplayLesson09 から呼び出されます
func GetComputerMoveDuringSelfPlay(kernel *e.Kernel, color e.Stone) e.Point {

	var start = time.Now()
	pl.AllPlayouts = 0

	var z, winRate = pl.GetBestZByUct(
		kernel,
		color,
		createPrintingOfCalc(),
		createPrintingOfCalcFin())

	var sec = time.Since(start).Seconds()
	code.Console.Info("(GetComputerMoveDuringSelfPlay) %.1f sec, %.0f playout/sec, play_z=%04d,rate=%.4f,movesNum=%d,color=%d,playouts=%d\n",
		sec, float64(pl.AllPlayouts)/sec, kernel.BoardCoordinate.GetZ4FromPoint(z), winRate, kernel.Position.MovesNum, color, pl.AllPlayouts)
	return z
}
