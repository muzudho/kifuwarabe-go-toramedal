package kernel

import "fmt"

// GetGtpMoveFromPoint - 番地をGTP用の指し手に変換。 例: Q10
func GetGtpMoveFromPoint(point Point) string {
	if point == 0 {
		return "PASS"
	} else if point == Cell_Illegal {
		return "ILLEGAL" // GTP の仕様外です
	}

	var y = int(point) / GetMemoryBoardWidth()
	var x = int(point) % GetMemoryBoardWidth()

	// 筋が25（'Z'）より大きくなることは想定していません
	var alphabet_x = 'A' + x - 1
	if alphabet_x >= 'I' {
		alphabet_x++
	}

	// code.Console.Debug("y=%d x=%d z=%d alphabet_x=%d alphabet_x(char)=%c\n", y, x, z, alphabet_x, alphabet_x)

	return fmt.Sprintf("%c%d", alphabet_x, y)
}
