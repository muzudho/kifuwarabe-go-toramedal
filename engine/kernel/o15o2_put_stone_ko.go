package kernel

// IsKo - コウに石を置いたか？
func (k *Kernel) IsPutStoneOnKo(point Point) bool {
	return k.Position.KoZ == point
}
