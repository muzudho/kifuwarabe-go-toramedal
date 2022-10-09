package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	code "github.com/muzudho/kifuwarabe-go-toramedal/engine/coding_obj"
	e "github.com/muzudho/kifuwarabe-go-toramedal/engine/kernel"
	pl "github.com/muzudho/kifuwarabe-go-toramedal/engine/play_algorithm"
	p "github.com/muzudho/kifuwarabe-go-toramedal/engine/presenter"
)

// RunGtpEngine - レッスン９a
// GTP2NNGS に対応しているのでは？
func RunGtpEngine(kernel *e.Kernel) {
	code.Console.Trace("# GoGo RunGtpEngine プログラム開始☆（＾～＾）\n")
	code.Console.Trace("# 何か標準入力しろだぜ☆（＾～＾）\n")

	// GUI から 囲碁エンジン へ入力があった、と考えてください
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var command = scanner.Text()
		code.Gtp.Log(command + "\n")
		code.ConsoleLog.Notice(command + "\n")

		var tokens = strings.Split(command, " ")
		switch tokens[0] {
		case "boardsize":
			// boardsize 19
			// 盤のサイズを変えます
			if 2 <= len(tokens) {
				var boardSize, err = strconv.Atoi(tokens[1])

				if err != nil {
					code.Console.Fatal(fmt.Sprintf("command=%s", command))
					panic(err)
				}

				kernel.BoardCoordinate.SetBoardSize(boardSize)
				e.SetBoardSize(boardSize)

				pl.InitKernel(kernel)

				code.Gtp.Print("= \n\n")
			} else {
				code.Gtp.Print("? unknown_command %s\n\n", command)
			}

		case "clear_board":
			pl.InitKernel(kernel)
			code.Gtp.Print("= \n\n")

		case "quit":
			os.Exit(0)

		case "protocol_version":
			code.Gtp.Print("= 2\n\n")

		case "name":
			code.Gtp.Print("= KifuwarabeMedal\n\n")

		case "version":
			code.Gtp.Print("= 0.0.3\n\n")

		case "list_commands":
			code.Gtp.Print("= boardsize\nclear_board\nquit\nprotocol_version\nundo\n" +
				"name\nversion\nlist_commands\nkomi\ngenmove\nplay\n\n")

		case "komi":
			// komi 6.5
			if 2 <= len(tokens) {
				var komi, err = strconv.ParseFloat(tokens[1], 64)

				if err != nil {
					code.Console.Fatal(fmt.Sprintf("command=%s", command))
					panic(err)
				}

				e.Komi = e.KomiFloat(komi)
				code.Gtp.Print("= %d\n\n", e.Komi)
			} else {
				code.Gtp.Print("? unknown_command %s\n\n", command)
			}

			// TODO 消す code.Gtp.Print("= 6.5\n\n")

		case "undo":
			// 未実装
			code.Gtp.Print("= \n\n")

		case "genmove":
			// genmove black
			// genmove white
			var color e.Stone
			if 1 < len(tokens) && strings.ToLower(tokens[1][0:1]) == "w" {
				color = 2
			} else {
				color = 1
			}
			var z = PlayComputerMoveLesson09a(kernel, color)
			code.Gtp.Print("= %s\n\n", e.GetGtpMoveFromPoint(z))

		case "play":
			// play black A3
			// play white D4
			// play black D5
			// play white E5
			// play black E4
			// play white D6
			// play black F5
			// play white C5
			// play black PASS
			// play white PASS
			if 2 < len(tokens) {
				var color e.Stone
				if strings.ToLower(tokens[1][0:1]) == "w" {
					color = 2
				} else {
					color = 1
				}

				var z = e.GetPointFromGtpMove(tokens[2])
				var recItem = new(e.RecordItem)
				recItem.Z = z
				recItem.Time = 0
				e.PutStoneOnRecord(kernel.Position, z, color, recItem)
				p.PrintBoard(kernel, kernel.Position.MovesNum)

				code.Gtp.Print("= \n\n")
			}

		default:
			code.Gtp.Print("? unknown_command\n\n")
		}
	}
}

// PlayComputerMoveLesson09a - コンピューター・プレイヤーの指し手。 SelfPlay, RunGtpEngine から呼び出されます。
func PlayComputerMoveLesson09a(
	kernel *e.Kernel,
	color e.Stone) e.Point {

	var st = time.Now()
	pl.AllPlayouts = 0

	var z, winRate = pl.GetBestZByUct(
		kernel,
		color,
		createPrintingOfCalc(),
		createPrintingOfCalcFin())

	if 1 < kernel.Position.MovesNum && // 初手ではないとして
		kernel.Position.Record[kernel.Position.MovesNum-1].GetZ() == 0 && // １つ前の手がパスで
		0.95 <= math.Abs(winRate) { // 95%以上の確率で勝ちか負けなら
		// こちらもパスします
		return 0
	}

	var sec = time.Since(st).Seconds()
	code.Console.Info("%.1f sec, %.0f playout/sec, play_z=%04d,rate=%.4f,movesNum=%d,color=%d,playouts=%d\n",
		sec, float64(pl.AllPlayouts)/sec, e.GetZ4FromPoint(z), winRate, kernel.Position.MovesNum, color, pl.AllPlayouts)

	var recItem = new(e.RecordItem)
	recItem.Z = z
	recItem.Time = sec
	e.PutStoneOnRecord(kernel.Position, z, color, recItem)
	p.PrintBoard(kernel, kernel.Position.MovesNum)

	return z
}
