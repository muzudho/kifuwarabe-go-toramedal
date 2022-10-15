package kernel

// RecordItem - 棋譜の1手分
type RecordItem struct {
	// 着手点
	placePlay Point
	// Time - 消費時間
	Time float64
	// コウ
	ko MutableKo
}

// SetPlacePlay - 着手点を設定
func (ri *RecordItem) SetPlacePlay(placePlay Point) {
	ri.placePlay = placePlay
}

// GetPlacePlay - 着手点を取得
func (ri *RecordItem) GetPlacePlay() Point {
	return ri.placePlay
}

// SetTime - 消費時間を設定
func (ri *RecordItem) SetTime(time float64) {
	ri.Time = time
}

// GetTime - 消費時間を取得
func (ri *RecordItem) GetTime() float64 {
	return ri.Time
}

// GetPlaceKo - コウの番地を取得
func (ri *RecordItem) GetPlaceKo() Point {
	return ri.ko.place
}

// SetPlaceKo - コウの番地を設定
func (ri *RecordItem) SetPlaceKo(placeKo Point) {
	ri.ko.place = placeKo
}

// SetPlaceKo - コウの番地を設定
func (ri *RecordItem) ClearPlaceKo() {
	ri.ko.ClearPlace()
}
