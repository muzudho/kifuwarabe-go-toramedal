package kernel

// GetStoneAtXy - 指定した交点の石の色
func (k *Kernel) GetStoneAtXy(x int, y int) Stone {
	var point = Point((y+1)*k.BoardCoordinate.GetMemoryBoardWidth() + x + 1)
	return k.Position.board.GetStoneAt(point)
}
