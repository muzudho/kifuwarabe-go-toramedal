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

	var boardWidth = kernel.Position.GetBoard().GetCoordinate().GetWidth()
	var boardHeight = kernel.Position.GetBoard().GetCoordinate().GetHeight()

	// Header
	b.WriteString("\n   ")
	for x := 0; x < boardWidth; x++ {
		b.WriteString(labelOfColumns[x+1])
	}
	b.WriteString("\n  +")
	for x := 0; x < boardWidth; x++ {
		b.WriteString("--")
	}
	b.WriteString("-+\n")

	// Body
	for j := 0; j < boardHeight; j++ {
		b.WriteString(labelOfRows[j+1])
		b.WriteString("|")
		for i := 0; i < boardWidth; i++ {
			var z = kernel.Position.GetBoard().GetCoordinate().GetPointFromXy(i+1, j+1)

			var mark = kernel.Position.GetCheckBoard().GetMarkAt(z)

			b.WriteString(numberLabels[mark])
		}
		b.WriteString(" |\n")
	}

	// Footer
	b.WriteString("  +")
	for x := 0; x < boardWidth; x++ {
		b.WriteString("--")
	}
	b.WriteString("-+\n")

	code.Console.Print(b.String())
}
