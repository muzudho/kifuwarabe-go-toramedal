package kernel

// IsKo - コウに石を置いたか？
func (p *Position) IsPutStoneOnKo(point Point) bool {
	return p.KoZ == point
}
