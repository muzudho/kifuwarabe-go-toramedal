package presenter

import (
	"strings"

	code "github.com/muzudho/kifuwarabe-go-toramedal/engine/coding_obj"
	e "github.com/muzudho/kifuwarabe-go-toramedal/engine/kernel"
)

// " 0" - 空点
// " 1" - 黒石
var numberLabels = [2]string{" 0", " 1"}

// PrintCheckBoard - チェックボードを描画。
func PrintCheckBoard(kernel *e.Kernel) {

	var b = &strings.Builder{}
	b.Grow(sz8k)

	var boardSize = e.BoardSize

	// Header
	b.WriteString("\n   ")
	for x := 0; x < boardSize; x++ {
		b.WriteString(labelOfColumns[x+1])
	}
	b.WriteString("\n  +")
	for x := 0; x < boardSize; x++ {
		b.WriteString("--")
	}
	b.WriteString("-+\n")

	// Body
	for y := 0; y < boardSize; y++ {
		b.WriteString(labelOfRows[y+1])
		b.WriteString("|")
		for x := 0; x < boardSize; x++ {
			var z = kernel.BoardCoordinate.GetPointFromXy(x, y)
			var number = kernel.Position.CheckAt(z)
			b.WriteString(numberLabels[number])
		}
		b.WriteString(" |\n")
	}

	// Footer
	b.WriteString("  +")
	for x := 0; x < boardSize; x++ {
		b.WriteString("--")
	}
	b.WriteString("-+\n")

	code.Console.Print(b.String())
}
