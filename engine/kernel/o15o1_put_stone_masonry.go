package kernel

// IsMasonry - 石の上に石を置こうとしたか？
func (b *Board) IsMasonry(point Point) bool {
	return !b.IsSpaceAt(point)
}
