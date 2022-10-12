package kernel

// Record - 棋譜
type Record struct {
	// 棋譜項目の配列
	items []RecordItem
	// コウの番地
	placeKo Point
}

// NewRecord - 新規作成
func NewRecord(maxPositionNumber PositionNumberInt) *Record {
	var r = new(Record)

	// 動的に長さがきまる配列を生成、その内容をインスタンスで埋めます
	r.items = make([]RecordItem, maxPositionNumber)
	for i := PositionNumberInt(0); i < maxPositionNumber; i++ {
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

// GetPlaceKo - コウの番地を取得
func (r *Record) GetPlaceKo() Point {
	return r.placeKo
}

// SetPlaceKo - コウの番地を設定
func (r *Record) SetPlaceKo(placeKo Point) {
	r.placeKo = placeKo
}
