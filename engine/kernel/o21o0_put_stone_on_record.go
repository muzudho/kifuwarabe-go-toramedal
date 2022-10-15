package kernel

import (
	"os"

	code "github.com/muzudho/kifuwarabe-go-toramedal/engine/coding_obj"
)

// PutStoneOnRecord - SelfPlay, RunGtpEngine から呼び出されます
func PutStoneOnRecord(kernel *Kernel, z Point, color Stone, recItem *RecordItem) {
	var err = PutStone(kernel, z, color)
	if err != 0 {
		code.Console.Error("(PutStoneOnRecord) Err!\n")
		os.Exit(0)
	}

	// 棋譜に記録
	kernel.Record.SetItemAt(kernel.Record.positionNumber, recItem)
	kernel.Record.positionNumber++
}
