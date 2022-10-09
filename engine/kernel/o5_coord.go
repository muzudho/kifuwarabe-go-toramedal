package kernel

import (
	"fmt"
	"strings"
)

// GetGtpMoveFromPoint - 番地をGTP用の指し手に変換。 例: Q10
func GetGtpMoveFromPoint(point Point) string {
	if point == 0 {
		return "PASS"
	} else if point == Cell_Illegal {
		return "PASS" // 仕方なく
		// return "ILLEGAL" // GTP の仕様外です
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

// GetPointFromXy - x,y 形式の座標を、 z （配列のインデックス）へ変換します。
// x,y は壁を含まない領域での 0 から始まる座標です。 z は壁を含む盤上での座標です
func GetPointFromXy(x int, y int) Point {
	return Point((y+1)*GetMemoryBoardWidth() + x + 1)
}

// GetPointFromGtpMove - GTPの座標符号を Point に変換します
// * `gtp_z` - 最初の１文字はアルファベット、２文字目（あれば３文字目）は数字と想定。 例: q10
func GetPointFromGtpMove(gtp_z string) Point {
	gtp_z = strings.ToUpper(gtp_z)

	if gtp_z == "PASS" {
		return 0
	}

	// 筋
	var x = gtp_z[0] - 'A' + 1
	if gtp_z[0] >= 'I' {
		x--
	}

	// 段
	var y = int(gtp_z[1] - '0')
	if 2 < len(gtp_z) {
		y *= 10
		y += int(gtp_z[2] - '0')
	}

	// インデックス
	var z = GetPointFromXy(int(x)-1, y-1)
	return z
}

// GetZ4FromPoint - z（配列のインデックス）を XXYY形式（3～4桁の数）の座標へ変換します。
func GetZ4FromPoint(point Point) int {
	if point == 0 {
		return 0
	}
	var y = int(point) / GetMemoryBoardWidth()
	var x = int(point) - y*GetMemoryBoardWidth()
	return x*100 + y
}
