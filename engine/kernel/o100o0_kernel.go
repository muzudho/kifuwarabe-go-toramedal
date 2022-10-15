package kernel

type Kernel struct {
	// Position - 局面
	Position *Position
	// Record - 棋譜
	Record Record
}

// NewKernel - 新規
func NewKernel(gameRule GameRule, boardSize int) *Kernel {
	var k = new(Kernel)

	k.Position = NewDirtyPosition(gameRule, boardSize)
	k.Record = *NewRecord(gameRule.GetMaxPositionNumber())

	return k
}

// ResizeBoard - 盤サイズの変更
// - 別途、盤面の初期化を行ってください
func (k *Kernel) ResizeBoard(boardSize int) {
	k.Position = NewDirtyPosition(k.Position.board.gameRule, boardSize)
}

// SetPlaceKoOfCurrentPosition - 現局面のコウの番地を設定
//
// - Tips 取った石が１個ならコウを疑う
func (k *Kernel) SetPlaceKoOfCurrentPosition(placeKo Point) {
	var posNum = k.Position.Number
	k.Record.items[posNum].SetPlaceKo(placeKo)
}

// ClearPlaceKoOfCurrentPosition - 現局面のコウの番地を消去
func (k *Kernel) ClearPlaceKoOfCurrentPosition() {
	var posNum = k.Position.Number
	k.Record.items[posNum].ClearPlaceKo()
}

// SetPlaceKo - 現局面のコウを取得
func (k *Kernel) GetPlaceKoOfCurrentPosition() Point {
	var posNum = k.Position.Number
	return k.Record.items[posNum].GetPlaceKo()
}
