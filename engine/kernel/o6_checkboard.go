package kernel

// Mark - 目印
type Mark uint8

const (
	Mark_Empty Mark = iota
	Mark_Checked
)

// CheckBoard - ２値ではなく多値
type CheckBoard struct {
	// 呼吸点を数えるための一時盤
	cells []Mark
}

func NewCheckBoard(memoryBoardArea int) *CheckBoard {
	var cb = new(CheckBoard)

	cb.cells = make([]Mark, memoryBoardArea)

	return cb
}

// GetMarkAt - 指定した交点の目印を取得
func (cb *CheckBoard) GetMarkAt(point Point) Mark {
	return cb.cells[point]
}

// SetCheckedAt - 指定した交点に目印を設定
func (cb *CheckBoard) SetMarkAt(point Point, mark Mark) {
	cb.cells[point] = mark
}

// IsEmptyAt - 指定した交点に目印は付いていないか？
func (cb *CheckBoard) IsEmptyAt(point Point) bool {
	return cb.cells[point] == Mark_Empty
}
