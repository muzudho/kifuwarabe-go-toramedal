package kernel

// Ko - コウ
type Ko struct {
	// コウの番地。無ければ Cell_Pass と同値
	place Point
}

// GetPlaceKo - コウの番地を取得
func (k *Ko) GetPlace() Point {
	return k.place
}

// SetPlaceKo - コウの番地を設定
func (k *Ko) SetPlace(placeKo Point) {
	k.place = placeKo
}
