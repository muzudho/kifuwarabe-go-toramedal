package presenter

import (
	"fmt"
	"strings"

	code "github.com/muzudho/kifuwarabe-go-toramedal/engine/coding_obj"
	e "github.com/muzudho/kifuwarabe-go-toramedal/engine/kernel"
)

var sz8k = 8 * 1024

// 案
//     A B C D E F G H J K L M N O P Q R S T
//   +---------------------------------------+
//  1| . . . . . . . . . . . . . . . . . . . |
//  2| . . . . . . . . . . . . . . . . . . . |
//  3| . . . . . . . . . . . . . . . . x . . |
//  4| . . . . . . . . . . . . . . . . . . . |
//  5| . . . . . . . . . . . . . . . . . . . |
//  6| . . . . . . . . . . . . . . . . . . . |
//  7| . . . . . . . . . . . . . . . . . . . |
//  8| . . . . . . . . . . . . . . . . . . . |
//  9| . . . . . . . . . . . . . . . . . . . |
// 10| . . . . . . . . . . . . . . . . . . . |
// 11| . . . . . . . . . . . . . . . . . . . |
// 12| . . . . . . . . . . . . . . . . . . . |
// 13| . . . . . . . . . . . . . . . . . . . |
// 14| . . . . . . . . . . . . . . . . . . . |
// 15| . . . . . . . . . . . . . . . . . . . |
// 16| . . . . . . . . . . . . . . . . . . . |
// 17| . . o . . . . . . . . . . . . . . . . |
// 18| . . . . . . . . . . . . . . . . . . . |
// 19| . . . . . . . . . . . . . . . . . . . |
//   +---------------------------------------+
//  Ko=0,positionNumber=999
//
// ASCII文字を使います（全角、半角の狂いがないため）
// 黒石は x 、 白石は o （ダークモードでもライトモードでも識別できるため）

// labelOfColumns - 各列の表示符号。
// 国際囲碁連盟のフォーマット
var labelOfColumns = [20]string{"xx", " A", " B", " C", " D", " E", " F", " G", " H", " J",
	" K", " L", " M", " N", " O", " P", " Q", " R", " S", " T"}

// labelOfRows - 各行の表示符号。
var labelOfRows = [20]string{" 0", " 1", " 2", " 3", " 4", " 5", " 6", " 7", " 8", " 9",
	"10", "11", "12", "13", "14", "15", "16", "17", "18", "19"}

// " ." - 空点
// " x" - 黒石
// " o" - 白石
// " #" - 枠（バグ目視確認用）
var stoneLabels = [4]string{" .", " x", " o", " #"}

// " ." - 空点（バグ目視確認用）
// " x" - 黒石（バグ目視確認用）
// " o" - 白石（バグ目視確認用）
// "+-" - 枠
var leftCornerLabels = [4]string{".", "x", "o", "+"}
var horizontalEdgeLabels = [4]string{" .", " x", " o", "--"}
var rightCornerLabels = [4]string{" .", " x", " o", "-+"}
var leftVerticalEdgeLabels = [4]string{".", "x", "o", "|"}
var rightVerticalEdgeLabels = [4]string{" .", " x", " o", " |"}

// PrintBoard - 盤を描画。
func PrintBoard(kernel *e.Kernel) {

	var sb = &strings.Builder{}
	sb.Grow(sz8k)

	if kernel.Position.GetBoard().GetGameRule().GetMaxPositionNumber() < kernel.Record.GetPositionNumber() {
		sb.WriteString(fmt.Sprintf("Out of bounds max position number %d.\r", kernel.Record.GetPositionNumber()))
	} else {
		var boardWidth = kernel.Position.GetBoard().GetCoordinate().GetWidth()
		var boardHeight = kernel.Position.GetBoard().GetCoordinate().GetHeight()

		// Header (numbers)
		sb.WriteString("\n   ")
		for x := 0; x < boardWidth; x++ {
			sb.WriteString(labelOfColumns[x+1])
		}
		// Header (line)
		sb.WriteString("\n  ")                                                     // number space
		sb.WriteString(leftCornerLabels[kernel.Position.GetBoard().GetStoneAt(0)]) // +
		for x := 0; x < boardWidth; x++ {
			sb.WriteString(horizontalEdgeLabels[kernel.Position.GetBoard().GetStoneAt(e.Point(x+1))]) // --
		}
		sb.WriteString(rightCornerLabels[kernel.Position.GetBoard().GetStoneAt(e.Point(kernel.Position.GetBoard().GetCoordinate().GetMemoryWidth()-1))]) // -+
		sb.WriteString("\n")

		// Body
		for y := 0; y < boardHeight; y++ {
			sb.WriteString(labelOfRows[y+1])                                                                                                                          // number
			sb.WriteString(leftVerticalEdgeLabels[kernel.Position.GetBoard().GetStoneAt(e.Point((y+1)*kernel.Position.GetBoard().GetCoordinate().GetMemoryWidth()))]) // |
			for x := 0; x < boardWidth; x++ {
				sb.WriteString(stoneLabels[kernel.GetStoneAtXy(x, y)])
			}
			sb.WriteString(rightVerticalEdgeLabels[kernel.Position.GetBoard().GetStoneAt(e.Point((y+2)*kernel.Position.GetBoard().GetCoordinate().GetMemoryWidth()-1))]) // " |"
			sb.WriteString("\n")
		}

		// Footer line
		sb.WriteString("  ") // number space
		var leftBottomCellNum = kernel.Position.GetBoard().GetCoordinate().GetMemoryWidth() * (kernel.Position.GetBoard().GetCoordinate().GetMemoryHeight() - 1)
		// 枠付きの盤なので、左下隅には `+` 石がある
		sb.WriteString(leftCornerLabels[kernel.Position.GetBoard().GetStoneAt(e.Point(leftBottomCellNum))]) // +
		for x := 0; x < boardWidth; x++ {
			sb.WriteString(horizontalEdgeLabels[kernel.Position.GetBoard().GetStoneAt(e.Point(leftBottomCellNum+x+1))]) // --
		}
		sb.WriteString(rightCornerLabels[kernel.Position.GetBoard().GetStoneAt(e.Point(kernel.Position.GetBoard().GetCoordinate().GetMemoryArea()-1))]) // -+
		sb.WriteString("\n")

		// Info
		sb.WriteString("  Ko=")
		if kernel.GetPlaceKoOfCurrentPosition() == e.Cell_Pass {
			sb.WriteString("_")
		} else {
			sb.WriteString(kernel.Position.GetBoard().GetCoordinate().GetGtpMoveFromPoint(kernel.GetPlaceKoOfCurrentPosition()))
		}
		sb.WriteString(fmt.Sprintf(",positionNumber=%d", kernel.Record.GetPositionNumber()))
		sb.WriteString("\n")
	}

	code.Console.Print(sb.String())
}
