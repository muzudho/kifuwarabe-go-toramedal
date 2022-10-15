package kernel

// Mark - 目印
type Mark uint8

const (
	Mark_Empty Mark = iota
	Mark_Checked
)

// CheckBoard - ２値ではなく多値
//
// - 呼吸点を数えるための一時盤
type CheckBoard struct {
	// 盤座標
	coordinate BoardCoordinate

	// 長さが可変な盤
	//
	// * 英語で交点は node かも知れないが、表計算でよく使われる cell の方を使う
	cells []Mark
}

// NewDirtyCheckBoard - 新規作成するが、初期化されていない
//
// * このメソッドを呼び出した後に Init 関数を呼び出してほしい
func NewDirtyCheckBoard() *CheckBoard {
	var cb = new(CheckBoard)

	cb.coordinate = BoardCoordinate{}

	return cb
}

// Init - 初期化
func (cb *CheckBoard) Init(newBoardCoordinate BoardCoordinate) {
	cb.coordinate = newBoardCoordinate
	cb.cells = make([]Mark, cb.coordinate.GetMemoryArea())
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
