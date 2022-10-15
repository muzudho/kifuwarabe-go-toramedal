package kernel

// IsMasonry - 石の上に石を置こうとしたか？
func (b *Board) IsMasonry(point Point) bool {
	// 空点以外に石を置こうとしたら、石の上に石を置いた扱いにする
	return !b.IsSpaceAt(point)
}
