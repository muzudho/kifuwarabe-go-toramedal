package presenter

import (
	code "github.com/muzudho/kifuwarabe-go-toramedal/engine/coding_obj"
	e "github.com/muzudho/kifuwarabe-go-toramedal/engine/kernel"
)

// PrintSgf - SGF形式の棋譜表示。
func PrintSgf(kernel *e.Kernel, positionNumber e.PositionNumberInt, record []*e.RecordItem) {
	var boardSize = kernel.Position.GetBoard().GetCoordinate().GetBoardWidth()

	code.Console.Print("(;GM[1]SZ[%d]KM[%.1f]PB[]PW[]\n", boardSize, kernel.Position.GetBoard().GetGameRule().GetKomi())
	for i := e.PositionNumberInt(0); i < positionNumber; i++ {
		var z = record[i].GetZ()
		var y = int(z) / kernel.Position.GetBoard().GetCoordinate().GetMemoryBoardWidth()
		var x = int(z) - y*kernel.Position.GetBoard().GetCoordinate().GetMemoryBoardWidth()
		var sStone = [2]string{"B", "W"}
		code.Console.Print(";%s", sStone[i&1])
		if z == 0 {
			code.Console.Print("[]")
		} else {
			code.Console.Print("[%c%c]", x+'a'-1, y+'a'-1)
		}
		if ((i + 1) % 10) == 0 {
			code.Console.Print("\n")
		}
	}
	code.Console.Print(")\n")
}
