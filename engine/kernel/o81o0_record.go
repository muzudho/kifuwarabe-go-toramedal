package kernel

// Record - 棋譜
type Record struct {
	// 棋譜項目の配列
	items []RecordItem
}

// NewRecord - 新規作成
func NewRecord(maxPositionNumber PositionNumberInt) *Record {
	var r = new(Record)

	// 動的に長さがきまる配列を生成、その内容をインスタンスで埋めます
	// 例えば、0手目が初期局面として、 400 手目まであるとすると、要素数は401要る。だから 1 足す
	var elementCount = maxPositionNumber + 1
	r.items = make([]RecordItem, elementCount)

	for i := PositionNumberInt(0); i < elementCount; i++ {
		r.items[i] = RecordItem{}
	}

	return r
}

func (r *Record) SetItemAt(posNum PositionNumberInt, item *RecordItem) {
	r.items[posNum] = *item
}

// GetPlacePlayAt - 指定局面の着手点を取得
func (r *Record) GetPlacePlayAt(posNum PositionNumberInt) Point {
	return r.items[posNum].placePlay
}

// SetPlacePlayAt - 指定局面の着手点を設定
func (r *Record) SetPlacePlayAt(posNum PositionNumberInt, placePlay Point) {
	r.items[posNum].placePlay = placePlay
}

// GetTimeAt - 指定局面の消費時間を取得
func (r *Record) GetTimeAt(posNum PositionNumberInt) float64 {
	return r.items[posNum].Time
}

// SetTimeAt - 指定局面の消費時間を設定
func (r *Record) SetTimeAt(posNum PositionNumberInt, time float64) {
	r.items[posNum].Time = time
}
