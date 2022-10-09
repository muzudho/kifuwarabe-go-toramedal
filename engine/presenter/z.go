package presenter

import (
	"strings"

	e "github.com/muzudho/kifuwarabe-go-toramedal/engine/kernel"
)

// GetZFromGtp - GTPの座標符号を z に変換します
// * `gtp_z` - 最初の１文字はアルファベット、２文字目（あれば３文字目）は数字と想定。 例: q10
func GetZFromGtp(position *e.Position, gtp_z string) e.Point {
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
	var z = e.GetPointFromXy(int(x)-1, y-1)
	// code.Console.Trace("# x=%d y=%d z=%d z4=%04d\n", x, y, z, position.GetZ4(z))
	return z
}
